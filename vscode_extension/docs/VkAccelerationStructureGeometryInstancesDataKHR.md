# VkAccelerationStructureGeometryInstancesDataKHR(3) Manual Page

## Name

VkAccelerationStructureGeometryInstancesDataKHR - Structure specifying a geometry consisting of instances of other acceleration structures



## [](#_c_specification)C Specification

The `VkAccelerationStructureGeometryInstancesDataKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureGeometryInstancesDataKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    VkBool32                         arrayOfPointers;
    VkDeviceOrHostAddressConstKHR    data;
} VkAccelerationStructureGeometryInstancesDataKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `arrayOfPointers` specifies whether `data` is used as an array of addresses or just an array.
- `data` is either the address of an array of device or host addresses referencing individual [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html) structures or packed motion instance information as described in [motion instances](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-motion-instances) if `arrayOfPointers` is `VK_TRUE`, or the address of an array of [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html) or [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceNV.html) structures. Addresses and [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html) structures are tightly packed. [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceNV.html) structures have a stride of 160 bytes.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureGeometryInstancesDataKHR-sType-sType)VUID-VkAccelerationStructureGeometryInstancesDataKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_INSTANCES_DATA_KHR`
- [](#VUID-VkAccelerationStructureGeometryInstancesDataKHR-pNext-pNext)VUID-VkAccelerationStructureGeometryInstancesDataKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryDataKHR.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureGeometryInstancesDataKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0