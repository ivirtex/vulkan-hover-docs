# vkGetAndroidHardwareBufferPropertiesANDROID(3) Manual Page

## Name

vkGetAndroidHardwareBufferPropertiesANDROID - Get Properties of External
Memory Android Hardware Buffers



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the memory parameters to use when importing an Android
hardware buffer, call:

``` c
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
VkResult vkGetAndroidHardwareBufferPropertiesANDROID(
    VkDevice                                    device,
    const struct AHardwareBuffer*               buffer,
    VkAndroidHardwareBufferPropertiesANDROID*   pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that will be importing `buffer`.

- `buffer` is the Android hardware buffer which will be imported.

- `pProperties` is a pointer to a
  [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferPropertiesANDROID.html)
  structure in which the properties of `buffer` are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-01884"
  id="VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-01884"></a>
  VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-01884  
  `buffer` **must** be a valid Android hardware buffer object with at
  least one of the `AHARDWAREBUFFER_USAGE_GPU_*` flags in its
  `AHardwareBuffer_Desc`::`usage`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetAndroidHardwareBufferPropertiesANDROID-device-parameter"
  id="VUID-vkGetAndroidHardwareBufferPropertiesANDROID-device-parameter"></a>
  VUID-vkGetAndroidHardwareBufferPropertiesANDROID-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-parameter"
  id="VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-parameter"></a>
  VUID-vkGetAndroidHardwareBufferPropertiesANDROID-buffer-parameter  
  `buffer` **must** be a valid pointer to a valid
  [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/AHardwareBuffer.html) value

- <a
  href="#VUID-vkGetAndroidHardwareBufferPropertiesANDROID-pProperties-parameter"
  id="VUID-vkGetAndroidHardwareBufferPropertiesANDROID-pProperties-parameter"></a>
  VUID-vkGetAndroidHardwareBufferPropertiesANDROID-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a
  [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferPropertiesANDROID.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html),
[VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferPropertiesANDROID.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetAndroidHardwareBufferPropertiesANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
