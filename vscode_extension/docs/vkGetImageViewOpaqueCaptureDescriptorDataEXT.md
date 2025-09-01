# vkGetImageViewOpaqueCaptureDescriptorDataEXT(3) Manual Page

## Name

vkGetImageViewOpaqueCaptureDescriptorDataEXT - Get image view opaque capture descriptor data



## [](#_c_specification)C Specification

To get the opaque capture descriptor data for an image view, call:

```c++
// Provided by VK_EXT_descriptor_buffer
VkResult vkGetImageViewOpaqueCaptureDescriptorDataEXT(
    VkDevice                                    device,
    const VkImageViewCaptureDescriptorDataInfoEXT* pInfo,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the data.
- `pInfo` is a pointer to a [VkImageViewCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCaptureDescriptorDataInfoEXT.html) structure specifying the image view.
- `pData` is a pointer to an application-allocated buffer where the data will be written.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-None-08080)VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-None-08080  
  The [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-pData-08081)VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-pData-08081  
  `pData` **must** point to a buffer that is at least [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`imageViewCaptureReplayDescriptorDataSize` bytes in size
- [](#VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-device-08082)VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-device-08082  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-device-parameter)VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-pInfo-parameter)VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkImageViewCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCaptureDescriptorDataInfoEXT.html) structure
- [](#VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-pData-parameter)VUID-vkGetImageViewOpaqueCaptureDescriptorDataEXT-pData-parameter  
  `pData` **must** be a pointer value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageViewCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCaptureDescriptorDataInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageViewOpaqueCaptureDescriptorDataEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0