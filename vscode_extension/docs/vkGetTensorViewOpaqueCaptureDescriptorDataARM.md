# vkGetTensorViewOpaqueCaptureDescriptorDataARM(3) Manual Page

## Name

vkGetTensorViewOpaqueCaptureDescriptorDataARM - Get tensor view opaque capture descriptor data



## [](#_c_specification)C Specification

To get the opaque capture descriptor data for a tensor view, call:

```c++
// Provided by VK_EXT_descriptor_buffer with VK_ARM_tensors
VkResult vkGetTensorViewOpaqueCaptureDescriptorDataARM(
    VkDevice                                    device,
    const VkTensorViewCaptureDescriptorDataInfoARM* pInfo,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the data.
- `pInfo` is a pointer to a [VkTensorViewCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCaptureDescriptorDataInfoARM.html) structure specifying the tensor view.
- `pData` is a pointer to a user-allocated buffer where the data will be written.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-descriptorBufferCaptureReplay-09706)VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-descriptorBufferCaptureReplay-09706  
  The [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBuffer) and [`descriptorBufferTensorDescriptors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferTensorDescriptors) features **must** be enabled
- [](#VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-pData-09707)VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-pData-09707  
  `pData` **must** point to a buffer that is at least [VkPhysicalDeviceDescriptorBufferTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferTensorPropertiesARM.html)::`tensorViewCaptureReplayDescriptorDataSize` bytes in size
- [](#VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-device-09708)VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-device-09708  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-device-parameter)VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-pInfo-parameter)VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkTensorViewCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCaptureDescriptorDataInfoARM.html) structure
- [](#VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-pData-parameter)VUID-vkGetTensorViewOpaqueCaptureDescriptorDataARM-pData-parameter  
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

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkTensorViewCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCaptureDescriptorDataInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetTensorViewOpaqueCaptureDescriptorDataARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0