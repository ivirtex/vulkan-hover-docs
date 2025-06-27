# VkAccelerationStructureMotionInstanceNV(3) Manual Page

## Name

VkAccelerationStructureMotionInstanceNV - Structure specifying a single
acceleration structure motion instance for building into an acceleration
structure geometry



## <a href="#_c_specification" class="anchor"></a>C Specification

*Acceleration structure motion instances* **can** be built into
top-level acceleration structures. Each acceleration structure instance
is a separate entry in the top-level acceleration structure which
includes all the geometry of a bottom-level acceleration structure at a
transformed location including a type of motion and parameters to
determine the motion of the instance over time.

An acceleration structure motion instance is defined by the structure:

``` c
// Provided by VK_NV_ray_tracing_motion_blur
typedef struct VkAccelerationStructureMotionInstanceNV {
    VkAccelerationStructureMotionInstanceTypeNV     type;
    VkAccelerationStructureMotionInstanceFlagsNV    flags;
    VkAccelerationStructureMotionInstanceDataNV     data;
} VkAccelerationStructureMotionInstanceNV;
```

## <a href="#_members" class="anchor"></a>Members

- `type` is a
  [VkAccelerationStructureMotionInstanceTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceTypeNV.html)
  enumerant identifying which type of motion instance this is and which
  type of the union is valid.

- `flags` is currently unused, but is required to keep natural alignment
  of `data`.

- `data` is a
  [VkAccelerationStructureMotionInstanceDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceDataNV.html)
  containing motion instance data for this instance.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If writing this other than with a standard C compiler, note that the
final structure should be 152 bytes in size.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureMotionInstanceNV-type-parameter"
  id="VUID-VkAccelerationStructureMotionInstanceNV-type-parameter"></a>
  VUID-VkAccelerationStructureMotionInstanceNV-type-parameter  
  `type` **must** be a valid
  [VkAccelerationStructureMotionInstanceTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceTypeNV.html)
  value

- <a
  href="#VUID-VkAccelerationStructureMotionInstanceNV-flags-zerobitmask"
  id="VUID-VkAccelerationStructureMotionInstanceNV-flags-zerobitmask"></a>
  VUID-VkAccelerationStructureMotionInstanceNV-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkAccelerationStructureMotionInstanceNV-staticInstance-parameter"
  id="VUID-VkAccelerationStructureMotionInstanceNV-staticInstance-parameter"></a>
  VUID-VkAccelerationStructureMotionInstanceNV-staticInstance-parameter  
  If `type` is
  `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_STATIC_NV`, the
  `staticInstance` member of `data` **must** be a valid
  [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInstanceKHR.html)
  structure

- <a
  href="#VUID-VkAccelerationStructureMotionInstanceNV-matrixMotionInstance-parameter"
  id="VUID-VkAccelerationStructureMotionInstanceNV-matrixMotionInstance-parameter"></a>
  VUID-VkAccelerationStructureMotionInstanceNV-matrixMotionInstance-parameter  
  If `type` is
  `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_MATRIX_MOTION_NV`, the
  `matrixMotionInstance` member of `data` **must** be a valid
  [VkAccelerationStructureMatrixMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMatrixMotionInstanceNV.html)
  structure

- <a
  href="#VUID-VkAccelerationStructureMotionInstanceNV-srtMotionInstance-parameter"
  id="VUID-VkAccelerationStructureMotionInstanceNV-srtMotionInstance-parameter"></a>
  VUID-VkAccelerationStructureMotionInstanceNV-srtMotionInstance-parameter  
  If `type` is
  `VK_ACCELERATION_STRUCTURE_MOTION_INSTANCE_TYPE_SRT_MOTION_NV`, the
  `srtMotionInstance` member of `data` **must** be a valid
  [VkAccelerationStructureSRTMotionInstanceNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureSRTMotionInstanceNV.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing_motion_blur](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing_motion_blur.html),
[VkAccelerationStructureMotionInstanceDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceDataNV.html),
[VkAccelerationStructureMotionInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceFlagsNV.html),
[VkAccelerationStructureMotionInstanceTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceTypeNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureMotionInstanceNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
