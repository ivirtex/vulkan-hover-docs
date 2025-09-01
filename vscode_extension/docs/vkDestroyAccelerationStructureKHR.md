# vkDestroyAccelerationStructureKHR(3) Manual Page

## Name

vkDestroyAccelerationStructureKHR - Destroy an acceleration structure object



## [](#_c_specification)C Specification

To destroy an acceleration structure, call:

```c++
// Provided by VK_KHR_acceleration_structure
void vkDestroyAccelerationStructureKHR(
    VkDevice                                    device,
    VkAccelerationStructureKHR                  accelerationStructure,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the acceleration structure.
- `accelerationStructure` is the acceleration structure to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-08934)VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-08934  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled
- [](#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02442)VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02442  
  All submitted commands that refer to `accelerationStructure` **must** have completed execution
- [](#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02443)VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02443  
  If `VkAllocationCallbacks` were provided when `accelerationStructure` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02444)VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-02444  
  If no `VkAllocationCallbacks` were provided when `accelerationStructure` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyAccelerationStructureKHR-device-parameter)VUID-vkDestroyAccelerationStructureKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parameter)VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parameter  
  If `accelerationStructure` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `accelerationStructure` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-vkDestroyAccelerationStructureKHR-pAllocator-parameter)VUID-vkDestroyAccelerationStructureKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parent)VUID-vkDestroyAccelerationStructureKHR-accelerationStructure-parent  
  If `accelerationStructure` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `accelerationStructure` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyAccelerationStructureKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0