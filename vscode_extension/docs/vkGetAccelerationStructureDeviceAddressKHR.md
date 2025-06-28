# vkGetAccelerationStructureDeviceAddressKHR(3) Manual Page

## Name

vkGetAccelerationStructureDeviceAddressKHR - Query an address of an acceleration structure



## [](#_c_specification)C Specification

To query the 64-bit device address for an acceleration structure, call:

```c++
// Provided by VK_KHR_acceleration_structure
VkDeviceAddress vkGetAccelerationStructureDeviceAddressKHR(
    VkDevice                                    device,
    const VkAccelerationStructureDeviceAddressInfoKHR* pInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that the acceleration structure was created on.
- `pInfo` is a pointer to a [VkAccelerationStructureDeviceAddressInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureDeviceAddressInfoKHR.html) structure specifying the acceleration structure to retrieve an address for.

## [](#_description)Description

The 64-bit return value is an address of the acceleration structure, which can be used for device and shader operations that involve acceleration structures, such as ray traversal and acceleration structure building.

If the acceleration structure was created with a non-zero value of [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`deviceAddress`, the return value will be the same address.

If the acceleration structure was created with a `type` of `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`, the returned address **must** be consistent with the relative offset to other acceleration structures with `type` `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR` allocated with the same [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html). That is, the difference in returned addresses between the two **must** be the same as the difference in offsets provided at acceleration structure creation.

The returned address **must** be aligned to 256 bytes.

Note

The acceleration structure device address **may** be different from the buffer device address corresponding to the acceleration structure’s start offset in its storage buffer for acceleration structure types other than `VK_ACCELERATION_STRUCTURE_TYPE_GENERIC_KHR`.

Valid Usage

- [](#VUID-vkGetAccelerationStructureDeviceAddressKHR-accelerationStructure-08935)VUID-vkGetAccelerationStructureDeviceAddressKHR-accelerationStructure-08935  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled
- [](#VUID-vkGetAccelerationStructureDeviceAddressKHR-device-03504)VUID-vkGetAccelerationStructureDeviceAddressKHR-device-03504  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled
- [](#VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09541)VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09541  
  If the buffer on which `pInfo->accelerationStructure` was placed is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09542)VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-09542  
  The buffer on which `pInfo->accelerationStructure` was placed **must** have been created with the `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT` usage flag

Valid Usage (Implicit)

- [](#VUID-vkGetAccelerationStructureDeviceAddressKHR-device-parameter)VUID-vkGetAccelerationStructureDeviceAddressKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-parameter)VUID-vkGetAccelerationStructureDeviceAddressKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkAccelerationStructureDeviceAddressInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureDeviceAddressInfoKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureDeviceAddressInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureDeviceAddressInfoKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetAccelerationStructureDeviceAddressKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0