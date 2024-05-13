# VkAccelerationStructureMatrixMotionInstanceNV(3) Manual Page

## Name

VkAccelerationStructureMatrixMotionInstanceNV - Structure specifying a
single acceleration structure matrix motion instance for building into
an acceleration structure geometry



## <a href="#_c_specification" class="anchor"></a>C Specification

An acceleration structure matrix motion instance is defined by the
structure:

``` c
// Provided by VK_NV_ray_tracing_motion_blur
typedef struct VkAccelerationStructureMatrixMotionInstanceNV {
    VkTransformMatrixKHR          transformT0;
    VkTransformMatrixKHR          transformT1;
    uint32_t                      instanceCustomIndex:24;
    uint32_t                      mask:8;
    uint32_t                      instanceShaderBindingTableRecordOffset:24;
    VkGeometryInstanceFlagsKHR    flags:8;
    uint64_t                      accelerationStructureReference;
} VkAccelerationStructureMatrixMotionInstanceNV;
```

## <a href="#_members" class="anchor"></a>Members

- `transformT0` is a [VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTransformMatrixKHR.html)
  structure describing a transformation to be applied to the
  acceleration structure at time 0.

- `transformT1` is a [VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTransformMatrixKHR.html)
  structure describing a transformation to be applied to the
  acceleration structure at time 1.

- `instanceCustomIndex` is a 24-bit user-specified index value
  accessible to ray shaders in the `InstanceCustomIndexKHR` built-in.

- `mask` is an 8-bit visibility mask for the geometry. The instance
  **may** only be hit if `Cull Mask & instance.mask != 0`

- `instanceShaderBindingTableRecordOffset` is a 24-bit offset used in
  calculating the hit shader binding table index.

- `flags` is an 8-bit mask of
  [VkGeometryInstanceFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryInstanceFlagBitsKHR.html)
  values to apply to this instance.

- `accelerationStructureReference` is either:

  - a device address containing the value obtained from
    [vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureDeviceAddressKHR.html)
    or
    [vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureHandleNV.html)
    (used by device operations which reference acceleration structures)
    or,

  - a [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html)
    object (used by host operations which reference acceleration
    structures).

## <a href="#_description" class="anchor"></a>Description

The C language specification does not define the ordering of bit-fields,
but in practice, this struct produces the correct layout with existing
compilers. The intended bit pattern is for the following:

- `instanceCustomIndex` and `mask` occupy the same memory as if a single
  `uint32_t` was specified in their place

  - `instanceCustomIndex` occupies the 24 least significant bits of that
    memory

  - `mask` occupies the 8 most significant bits of that memory

- `instanceShaderBindingTableRecordOffset` and `flags` occupy the same
  memory as if a single `uint32_t` was specified in their place

  - `instanceShaderBindingTableRecordOffset` occupies the 24 least
    significant bits of that memory

  - `flags` occupies the 8 most significant bits of that memory

If a compiler produces code that diverges from that pattern,
applications **must** employ another method to set values according to
the correct bit pattern.

The transform for a matrix motion instance at a point in time is derived
by component-wise linear interpolation of the two transforms. That is,
for a `time` in \[0,1\] the resulting transform is

  
`transformT0` × (1 - `time`) + `transformT1` × `time`

Valid Usage (Implicit)

- <a
  href="#VUID-VkAccelerationStructureMatrixMotionInstanceNV-flags-parameter"
  id="VUID-VkAccelerationStructureMatrixMotionInstanceNV-flags-parameter"></a>
  VUID-VkAccelerationStructureMatrixMotionInstanceNV-flags-parameter  
  `flags` **must** be a valid combination of
  [VkGeometryInstanceFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryInstanceFlagBitsKHR.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_motion_blur](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_motion_blur.html),
[VkAccelerationStructureMotionInstanceDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceDataNV.html),
[VkGeometryInstanceFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryInstanceFlagsKHR.html),
[VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTransformMatrixKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureMatrixMotionInstanceNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
