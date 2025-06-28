# vkGetImageOpaqueCaptureDescriptorDataEXT(3) Manual Page

## Name

vkGetImageOpaqueCaptureDescriptorDataEXT - Get image opaque capture descriptor data



## [](#_c_specification)C Specification

To get the opaque capture descriptor data for an image, call:

```c++
// Provided by VK_EXT_descriptor_buffer
VkResult vkGetImageOpaqueCaptureDescriptorDataEXT(
    VkDevice                                    device,
    const VkImageCaptureDescriptorDataInfoEXT*  pInfo,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the data.
- `pInfo` is a pointer to a [VkImageCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCaptureDescriptorDataInfoEXT.html) structure specifying the image.
- `pData` is a pointer to an application-allocated buffer where the data will be written.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-None-08076)VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-None-08076  
  The [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-08077)VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-08077  
  `pData` **must** point to a buffer that is at least [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`imageCaptureReplayDescriptorDataSize` bytes in size
- [](#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-08078)VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-08078  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-parameter)VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pInfo-parameter)VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkImageCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCaptureDescriptorDataInfoEXT.html) structure
- [](#VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-parameter)VUID-vkGetImageOpaqueCaptureDescriptorDataEXT-pData-parameter  
  `pData` **must** be a pointer value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCaptureDescriptorDataInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageOpaqueCaptureDescriptorDataEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0