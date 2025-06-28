# vkGetTensorOpaqueCaptureDescriptorDataARM(3) Manual Page

## Name

vkGetTensorOpaqueCaptureDescriptorDataARM - Get tensor opaque capture descriptor data



## [](#_c_specification)C Specification

To get the opaque capture descriptor data for a tensor, call:

```c++
// Provided by VK_EXT_descriptor_buffer with VK_ARM_tensors
VkResult vkGetTensorOpaqueCaptureDescriptorDataARM(
    VkDevice                                    device,
    const VkTensorCaptureDescriptorDataInfoARM* pInfo,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the data.
- `pInfo` is a pointer to a [VkTensorCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCaptureDescriptorDataInfoARM.html) structure specifying the tensor.
- `pData` is a pointer to a user-allocated buffer where the data will be written.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-descriptorBufferCaptureReplay-09702)VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-descriptorBufferCaptureReplay-09702  
  The [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBuffer) and [pname::descriptorBufferTensorDescriptors](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferTensorDescriptors) features **must** be enabled
- [](#VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-pData-09703)VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-pData-09703  
  `pData` **must** point to a buffer that is at least [VkPhysicalDeviceDescriptorBufferTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferTensorPropertiesARM.html)::`tensorCaptureReplayDescriptorDataSize` bytes in size
- [](#VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-device-09704)VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-device-09704  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-device-parameter)VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-pInfo-parameter)VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkTensorCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCaptureDescriptorDataInfoARM.html) structure
- [](#VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-pData-parameter)VUID-vkGetTensorOpaqueCaptureDescriptorDataARM-pData-parameter  
  `pData` **must** be a pointer value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkTensorCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCaptureDescriptorDataInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetTensorOpaqueCaptureDescriptorDataARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0