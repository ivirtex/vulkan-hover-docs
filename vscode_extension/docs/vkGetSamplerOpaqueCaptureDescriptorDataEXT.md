# vkGetSamplerOpaqueCaptureDescriptorDataEXT(3) Manual Page

## Name

vkGetSamplerOpaqueCaptureDescriptorDataEXT - Get sampler opaque capture descriptor data



## [](#_c_specification)C Specification

To get the opaque capture descriptor data for a sampler, call:

```c++
// Provided by VK_EXT_descriptor_buffer
VkResult vkGetSamplerOpaqueCaptureDescriptorDataEXT(
    VkDevice                                    device,
    const VkSamplerCaptureDescriptorDataInfoEXT* pInfo,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the data.
- `pInfo` is a pointer to a [VkSamplerCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCaptureDescriptorDataInfoEXT.html) structure specifying the sampler.
- `pData` is a pointer to an application-allocated buffer where the data will be written.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-None-08084)VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-None-08084  
  The [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-08085)VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-08085  
  `pData` **must** point to a buffer that is at least [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`samplerCaptureReplayDescriptorDataSize` bytes in size
- [](#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-08086)VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-08086  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-parameter)VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pInfo-parameter)VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkSamplerCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCaptureDescriptorDataInfoEXT.html) structure
- [](#VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-parameter)VUID-vkGetSamplerOpaqueCaptureDescriptorDataEXT-pData-parameter  
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

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSamplerCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCaptureDescriptorDataInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetSamplerOpaqueCaptureDescriptorDataEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0