# VkPhysicalDeviceAccelerationStructurePropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceAccelerationStructurePropertiesKHR - Properties of the physical device for acceleration structure



## [](#_c_specification)C Specification

The `VkPhysicalDeviceAccelerationStructurePropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkPhysicalDeviceAccelerationStructurePropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           maxGeometryCount;
    uint64_t           maxInstanceCount;
    uint64_t           maxPrimitiveCount;
    uint32_t           maxPerStageDescriptorAccelerationStructures;
    uint32_t           maxPerStageDescriptorUpdateAfterBindAccelerationStructures;
    uint32_t           maxDescriptorSetAccelerationStructures;
    uint32_t           maxDescriptorSetUpdateAfterBindAccelerationStructures;
    uint32_t           minAccelerationStructureScratchOffsetAlignment;
} VkPhysicalDeviceAccelerationStructurePropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxGeometryCount` is the maximum number of geometries in the bottom level acceleration structure.
- []()`maxInstanceCount` is the maximum number of instances in the top level acceleration structure.
- []()`maxPrimitiveCount` is the maximum number of triangles or AABBs in all geometries in the bottom level acceleration structure.
- []()`maxPerStageDescriptorAccelerationStructures` is the maximum number of acceleration structure bindings that **can** be accessible to a single shader stage in a pipeline layout. Descriptor bindings with a descriptor type of `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` count against this limit. Only descriptor bindings in descriptor set layouts created without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set count against this limit.
- []()`maxPerStageDescriptorUpdateAfterBindAccelerationStructures` is similar to `maxPerStageDescriptorAccelerationStructures` but counts descriptor bindings from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetAccelerationStructures` is the maximum number of acceleration structure descriptors that **can** be included in descriptor bindings in a pipeline layout across all pipeline shader stages and descriptor set numbers. Descriptor bindings with a descriptor type of `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` count against this limit. Only descriptor bindings in descriptor set layouts created without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set count against this limit.
- []()`maxDescriptorSetUpdateAfterBindAccelerationStructures` is similar to `maxDescriptorSetAccelerationStructures` but counts descriptor bindings from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`minAccelerationStructureScratchOffsetAlignment` is the minimum **required** alignment, in bytes, for scratch data passed in to an acceleration structure build command. The value **must** be a power of two.

## [](#_description)Description

Due to the fact that the geometry, instance, and primitive counts are specified at acceleration structure creation as 32-bit values, [`maxGeometryCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxGeometryCount), [`maxInstanceCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxInstanceCount), and [`maxPrimitiveCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxPrimitiveCount) **must** not exceed 232-1.

If the `VkPhysicalDeviceAccelerationStructurePropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Limits specified by this structure **must** match those specified with the same name in [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceAccelerationStructurePropertiesKHR-sType-sType)VUID-VkPhysicalDeviceAccelerationStructurePropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceAccelerationStructurePropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0