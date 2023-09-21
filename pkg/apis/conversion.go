/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package apis

import (
	url "net/url"
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*corev1.NodeProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return ConvertURLValuesToV1NodeProxyOptions(a.(*url.Values), b.(*corev1.NodeProxyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*corev1.PodAttachOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return ConvertURLValuesToV1PodAttachOptions(a.(*url.Values), b.(*corev1.PodAttachOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*corev1.PodExecOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return ConvertURLValuesToV1PodExecOptions(a.(*url.Values), b.(*corev1.PodExecOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*corev1.PodLogOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return ConvertURLValuesToV1PodLogOptions(a.(*url.Values), b.(*corev1.PodLogOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*corev1.PodPortForwardOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return ConvertURLValuesToV1PodPortForwardOptions(a.(*url.Values), b.(*corev1.PodPortForwardOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*corev1.PodProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return ConvertURLValuesToV1PodProxyOptions(a.(*url.Values), b.(*corev1.PodProxyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*corev1.ServiceProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return ConvertURLValuesToV1ServiceProxyOptions(a.(*url.Values), b.(*corev1.ServiceProxyOptions), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvertURLValuesToV1NodeProxyOptions(in *url.Values, out *corev1.NodeProxyOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["path"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Path, s); err != nil {
			return err
		}
	} else {
		out.Path = ""
	}
	return nil
}

// ConvertURLValuesToV1NodeProxyOptions is an autogenerated conversion function.
func ConvertURLValuesToV1NodeProxyOptions(in *url.Values, out *corev1.NodeProxyOptions, s conversion.Scope) error {
	return autoConvertURLValuesToV1NodeProxyOptions(in, out, s)
}

func autoConvertURLValuesToV1PodAttachOptions(in *url.Values, out *corev1.PodAttachOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["stdin"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.Stdin, s); err != nil {
			return err
		}
	} else {
		out.Stdin = false
	}
	if values, ok := map[string][]string(*in)["stdout"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.Stdout, s); err != nil {
			return err
		}
	} else {
		out.Stdout = false
	}
	if values, ok := map[string][]string(*in)["stderr"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.Stderr, s); err != nil {
			return err
		}
	} else {
		out.Stderr = false
	}
	if values, ok := map[string][]string(*in)["tty"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.TTY, s); err != nil {
			return err
		}
	} else {
		out.TTY = false
	}
	if values, ok := map[string][]string(*in)["container"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Container, s); err != nil {
			return err
		}
	} else {
		out.Container = ""
	}
	return nil
}

// Convert_url_Values_To_v1_PodAttachOptions is an autogenerated conversion function.
func ConvertURLValuesToV1PodAttachOptions(in *url.Values, out *corev1.PodAttachOptions, s conversion.Scope) error {
	return autoConvertURLValuesToV1PodAttachOptions(in, out, s)
}

func autoConvertURLValuesToV1PodExecOptions(in *url.Values, out *corev1.PodExecOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["stdin"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.Stdin, s); err != nil {
			return err
		}
	} else {
		out.Stdin = false
	}
	if values, ok := map[string][]string(*in)["stdout"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.Stdout, s); err != nil {
			return err
		}
	} else {
		out.Stdout = false
	}
	if values, ok := map[string][]string(*in)["stderr"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.Stderr, s); err != nil {
			return err
		}
	} else {
		out.Stderr = false
	}
	if values, ok := map[string][]string(*in)["tty"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.TTY, s); err != nil {
			return err
		}
	} else {
		out.TTY = false
	}
	if values, ok := map[string][]string(*in)["container"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Container, s); err != nil {
			return err
		}
	} else {
		out.Container = ""
	}
	if values, ok := map[string][]string(*in)["command"]; ok && len(values) > 0 {
		out.Command = *(*[]string)(unsafe.Pointer(&values))
	} else {
		out.Command = nil
	}
	return nil
}

// ConvertURLValuesToV1PodExecOptions is an autogenerated conversion function.
func ConvertURLValuesToV1PodExecOptions(in *url.Values, out *corev1.PodExecOptions, s conversion.Scope) error {
	return autoConvertURLValuesToV1PodExecOptions(in, out, s)
}

func autoConvertURLValuesToV1PodLogOptions(in *url.Values, out *corev1.PodLogOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["container"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Container, s); err != nil {
			return err
		}
	} else {
		out.Container = ""
	}
	if values, ok := map[string][]string(*in)["follow"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.Follow, s); err != nil {
			return err
		}
	} else {
		out.Follow = false
	}
	if values, ok := map[string][]string(*in)["previous"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.Previous, s); err != nil {
			return err
		}
	} else {
		out.Previous = false
	}
	if values, ok := map[string][]string(*in)["sinceSeconds"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_Pointer_int64(&values, &out.SinceSeconds, s); err != nil {
			return err
		}
	} else {
		out.SinceSeconds = nil
	}
	if values, ok := map[string][]string(*in)["sinceTime"]; ok && len(values) > 0 {
		if err := metav1.Convert_Slice_string_To_Pointer_v1_Time(&values, &out.SinceTime, s); err != nil {
			return err
		}
	} else {
		out.SinceTime = nil
	}
	if values, ok := map[string][]string(*in)["timestamps"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.Timestamps, s); err != nil {
			return err
		}
	} else {
		out.Timestamps = false
	}
	if values, ok := map[string][]string(*in)["tailLines"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_Pointer_int64(&values, &out.TailLines, s); err != nil {
			return err
		}
	} else {
		out.TailLines = nil
	}
	if values, ok := map[string][]string(*in)["limitBytes"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_Pointer_int64(&values, &out.LimitBytes, s); err != nil {
			return err
		}
	} else {
		out.LimitBytes = nil
	}
	if values, ok := map[string][]string(*in)["insecureSkipTLSVerifyBackend"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_bool(&values, &out.InsecureSkipTLSVerifyBackend, s); err != nil {
			return err
		}
	} else {
		out.InsecureSkipTLSVerifyBackend = false
	}
	return nil
}

// ConvertURLValuesToV1PodLogOptions is an autogenerated conversion function.
func ConvertURLValuesToV1PodLogOptions(in *url.Values, out *corev1.PodLogOptions, s conversion.Scope) error {
	return autoConvertURLValuesToV1PodLogOptions(in, out, s)
}

func autoConvertURLValuesToV1PodPortForwardOptions(in *url.Values, out *corev1.PodPortForwardOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["ports"]; ok && len(values) > 0 {
		if err := metav1.Convert_Slice_string_To_Slice_int32(&values, &out.Ports, s); err != nil {
			return err
		}
	} else {
		out.Ports = nil
	}
	return nil
}

// ConvertURLValuesToV1PodPortForwardOptions is an autogenerated conversion function.
func ConvertURLValuesToV1PodPortForwardOptions(in *url.Values, out *corev1.PodPortForwardOptions, s conversion.Scope) error {
	return autoConvertURLValuesToV1PodPortForwardOptions(in, out, s)
}

func autoConvertURLValuesToV1PodProxyOptions(in *url.Values, out *corev1.PodProxyOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["path"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Path, s); err != nil {
			return err
		}
	} else {
		out.Path = ""
	}
	return nil
}

// ConvertURLValuesToV1PodProxyOptions is an autogenerated conversion function.
func ConvertURLValuesToV1PodProxyOptions(in *url.Values, out *corev1.PodProxyOptions, s conversion.Scope) error {
	return autoConvertURLValuesToV1PodProxyOptions(in, out, s)
}

func autoConvertURLValuesToV1ServiceProxyOptions(in *url.Values, out *corev1.ServiceProxyOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["path"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Path, s); err != nil {
			return err
		}
	} else {
		out.Path = ""
	}
	return nil
}

// ConvertURLValuesToV1ServiceProxyOptions is an autogenerated conversion function.
func ConvertURLValuesToV1ServiceProxyOptions(in *url.Values, out *corev1.ServiceProxyOptions, s conversion.Scope) error {
	return autoConvertURLValuesToV1ServiceProxyOptions(in, out, s)
}
