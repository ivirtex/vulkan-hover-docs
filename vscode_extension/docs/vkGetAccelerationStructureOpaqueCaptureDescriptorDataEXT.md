# vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT(3) Manual Page

## Name

vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT - Get acceleration structure opaque capture descriptor data



## [](#_c_specification)C Specification

To get the opaque capture descriptor data for an acceleration structure, call:

```c++
// Provided by VK_EXT_descriptor_buffer with VK_KHR_acceleration_structure or VK_NV_ray_tracing
VkResult vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT(
    VkDevice                                    device,
    const VkAccelerationStructureCaptureDescriptorDataInfoEXT* pInfo,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device that gets the data.
- `pInfo` is a pointer to a [VkAccelerationStructureCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCaptureDescriptorDataInfoEXT.html) structure specifying the acceleration structure.
- `pData` is a pointer to an application-allocated buffer where the data will be written.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-None-08088)VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-None-08088  
  The [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-08089)VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-08089  
  `pData` **must** point to a buffer that is at least [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`accelerationStructureCaptureReplayDescriptorDataSize` bytes in size
- [](#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-08090)VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-08090  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-parameter)VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pInfo-parameter)VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkAccelerationStructureCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCaptureDescriptorDataInfoEXT.html) structure
- [](#VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-parameter)VUID-vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT-pData-parameter  
  `pData` **must** be a pointer value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCaptureDescriptorDataInfoEXT.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0