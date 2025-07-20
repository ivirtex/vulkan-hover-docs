# vkGetAndroidHardwareBufferPropertiesANDROID(3) Manual Page

## Name

vkGetAndroidHardwareBufferPropertiesANDROID - Get Properties of External Memory Android Hardware Buffers



## [](#_c_specification)C Specification

To determine the memory parameters to use when importing an Android hardware buffer, call:

```c++
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
VkResult vkGetAndroidHardwareBufferPropertiesANDROID(
    VkDevice                                    device,
    const struct AHardwareBuffer*               buffer,
    VkAndroidHardwareBufferPropertiesANDROID*   pProperties);
```

## [](#_parameters)Parameters

- `device` is the logical device that will be importing `buffer`.
- `buffer` is the Android hardware buffer which will be imported.
- `pProperties` is a pointer to a [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferPropertiesANDROID.html) structure in which the properties of `buffer` are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-01884)VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-01884  
  `buffer` **must** be a valid Android hardware buffer object with at least one of the `AHARDWAREBUFFER_USAGE_GPU_*` flags in its `AHardwareBuffer_Desc`::`usage`

Valid Usage (Implicit)

- [](#VUID-vkGetAndroidHardwareBufferPropertiesANDROID-device-parameter)VUID-vkGetAndroidHardwareBufferPropertiesANDROID-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-parameter)VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-parameter  
  `buffer` **must** be a valid pointer to a valid [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) value
- [](#VUID-vkGetAndroidHardwareBufferPropertiesANDROID-pProperties-parameter)VUID-vkGetAndroidHardwareBufferPropertiesANDROID-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferPropertiesANDROID.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_ANDROID\_external\_memory\_android\_hardware\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html), [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferPropertiesANDROID.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetAndroidHardwareBufferPropertiesANDROID)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0