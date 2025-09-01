# VkAccelerationStructureInstanceKHR(3) Manual Page

## Name

VkAccelerationStructureInstanceKHR - Structure specifying a single acceleration structure instance for building into an acceleration structure geometry



## [](#_c_specification)C Specification

*Acceleration structure instances* **can** be built into top-level acceleration structures. Each acceleration structure instance is a separate entry in the top-level acceleration structure which includes all the geometry of a bottom-level acceleration structure at a transformed location. Multiple instances **can** point to the same bottom level acceleration structure.

An acceleration structure instance is defined by the structure:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureInstanceKHR {
    VkTransformMatrixKHR          transform;
    uint32_t                      instanceCustomIndex:24;
    uint32_t                      mask:8;
    uint32_t                      instanceShaderBindingTableRecordOffset:24;
    VkGeometryInstanceFlagsKHR    flags:8;
    uint64_t                      accelerationStructureReference;
} VkAccelerationStructureInstanceKHR;
```

or the equivalent

```c++
// Provided by VK_NV_ray_tracing
typedef VkAccelerationStructureInstanceKHR VkAccelerationStructureInstanceNV;
```

## [](#_members)Members

- `transform` is a [VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTransformMatrixKHR.html) structure describing a transformation to be applied to the acceleration structure.
- `instanceCustomIndex` is a 24-bit application-specified index value accessible to ray shaders in the `InstanceCustomIndexKHR` built-in.
- `mask` is an 8-bit visibility mask for the geometry. The instance **may** only be hit if `Cull Mask & instance.mask != 0`
- `instanceShaderBindingTableRecordOffset` is a 24-bit offset used in calculating the hit shader binding table index.
- `flags` is an 8-bit mask of [VkGeometryInstanceFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryInstanceFlagBitsKHR.html) values to apply to this instance.
- `accelerationStructureReference` is either :
  
  - a device address containing the value obtained from [vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureDeviceAddressKHR.html) or [vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureHandleNV.html) (used by device operations which reference acceleration structures) or,
  - a device address containing a bottom level cluster acceleration structure built using [vkCmdBuildClusterAccelerationStructureIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBuildClusterAccelerationStructureIndirectNV.html)
  - a [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) object (used by host operations which reference acceleration structures).

## [](#_description)Description

The C language specification does not define the ordering of bit-fields, but in practice, this structure produces the correct layout with existing compilers. The intended bit pattern is for the following:

- `instanceCustomIndex` and `mask` occupy the same memory as if a single `uint32_t` was specified in their place
  
  - `instanceCustomIndex` occupies the 24 least significant bits of that memory
  - `mask` occupies the 8 most significant bits of that memory
- `instanceShaderBindingTableRecordOffset` and `flags` occupy the same memory as if a single `uint32_t` was specified in their place
  
  - `instanceShaderBindingTableRecordOffset` occupies the 24 least significant bits of that memory
  - `flags` occupies the 8 most significant bits of that memory

If a compiler produces code that diverges from that pattern, applications **must** employ another method to set values according to the correct bit pattern.

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureInstanceKHR-flags-parameter)VUID-VkAccelerationStructureInstanceKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkGeometryInstanceFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryInstanceFlagBitsKHR.html) values

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureMotionInstanceDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInstanceDataNV.html), [VkGeometryInstanceFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryInstanceFlagsKHR.html), [VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTransformMatrixKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureInstanceKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0