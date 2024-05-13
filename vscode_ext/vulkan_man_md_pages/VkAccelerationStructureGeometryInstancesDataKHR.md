# VkAccelerationStructureGeometryInstancesDataKHR(3) Manual Page

## Name

VkAccelerationStructureGeometryInstancesDataKHR - Structure specifying a
geometry consisting of instances of other acceleration structures



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureGeometryInstancesDataKHR` structure is
defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureGeometryInstancesDataKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    VkBool32                         arrayOfPointers;
    VkDeviceOrHostAddressConstKHR    data;
} VkAccelerationStructureGeometryInstancesDataKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `arrayOfPointers` specifies whether `data` is used as an array of
  addresses or just an array.

- `data` is either the address of an array of device or host addresses
  referencing individual
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)
  structures or packed motion instance information as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-motion-instances"
  target="_blank" rel="noopener">motion instances</a> if
  `arrayOfPointers` is `VK_TRUE`, or the address of an array of
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)
  or
  [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceNV.html)
  structures. Addresses and
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)
  structures are tightly packed.
  [VkAccelerationStructureMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceNV.html)
  structures have a stride of 160 bytes.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkAccelerationStructureGeometryInstancesDataKHR-sType-sType"
  id="VUID-VkAccelerationStructureGeometryInstancesDataKHR-sType-sType"></a>
  VUID-VkAccelerationStructureGeometryInstancesDataKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_INSTANCES_DATA_KHR`

- <a
  href="#VUID-VkAccelerationStructureGeometryInstancesDataKHR-pNext-pNext"
  id="VUID-VkAccelerationStructureGeometryInstancesDataKHR-pNext-pNext"></a>
  VUID-VkAccelerationStructureGeometryInstancesDataKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryDataKHR.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureGeometryInstancesDataKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
