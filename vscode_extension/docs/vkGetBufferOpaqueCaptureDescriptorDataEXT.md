# vkGetBufferOpaqueCaptureDescriptorDataEXT(3) Manual Page

## Name

vkGetBufferOpaqueCaptureDescriptorDataEXT - Get buffer opaque capture descriptor data



## [](#_c_specification)C Specification

To get the opaque descriptor data for a buffer, call:

```c++
// Provided by VK_EXT_descriptor_buffer
VkResult vkGetBufferOpaqueCaptureDescriptorDataEXT(
    VkDevice                                    device,
    const VkBufferCaptureDescriptorDataInfoEXT* pInfo,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the data.
- `pInfo` is a pointer to a [VkBufferCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCaptureDescriptorDataInfoEXT.html) structure specifying the buffer.
- `pData` is a pointer to an application-allocated buffer where the data will be written.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-None-08072)VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-None-08072  
  The [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-08073)VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-08073  
  `pData` **must** point to a buffer that is at least [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`bufferCaptureReplayDescriptorDataSize` bytes in size
- [](#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-08074)VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-08074  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-parameter)VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pInfo-parameter)VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkBufferCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCaptureDescriptorDataInfoEXT.html) structure
- [](#VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-parameter)VUID-vkGetBufferOpaqueCaptureDescriptorDataEXT-pData-parameter  
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

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkBufferCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCaptureDescriptorDataInfoEXT.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetBufferOpaqueCaptureDescriptorDataEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0