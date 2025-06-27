# VkCopyAccelerationStructureInfoKHR(3) Manual Page

## Name

VkCopyAccelerationStructureInfoKHR - Parameters for copying an
acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyAccelerationStructureInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkCopyAccelerationStructureInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkAccelerationStructureKHR            src;
    VkAccelerationStructureKHR            dst;
    VkCopyAccelerationStructureModeKHR    mode;
} VkCopyAccelerationStructureInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `src` is the source acceleration structure for the copy.

- `dst` is the target acceleration structure for the copy.

- `mode` is a
  [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html)
  value specifying additional operations to perform during the copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-mode-03410"
  id="VUID-VkCopyAccelerationStructureInfoKHR-mode-03410"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-mode-03410  
  `mode` **must** be `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR`
  or `VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR`

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-src-04963"
  id="VUID-VkCopyAccelerationStructureInfoKHR-src-04963"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-src-04963  
  The source acceleration structure `src` **must** have been constructed
  prior to the execution of this command

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-src-03411"
  id="VUID-VkCopyAccelerationStructureInfoKHR-src-03411"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-src-03411  
  If `mode` is `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR`, `src`
  **must** have been constructed with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` in the
  build

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-buffer-03718"
  id="VUID-VkCopyAccelerationStructureInfoKHR-buffer-03718"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-buffer-03718  
  The `buffer` used to create `src` **must** be bound to device memory

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-buffer-03719"
  id="VUID-VkCopyAccelerationStructureInfoKHR-buffer-03719"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-buffer-03719  
  The `buffer` used to create `dst` **must** be bound to device memory

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-dst-07791"
  id="VUID-VkCopyAccelerationStructureInfoKHR-dst-07791"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-dst-07791  
  The range of memory backing `dst` that is accessed by this command
  **must** not overlap the memory backing `src` that is accessed by this
  command

Valid Usage (Implicit)

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-sType-sType"
  id="VUID-VkCopyAccelerationStructureInfoKHR-sType-sType"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_INFO_KHR`

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-pNext-pNext"
  id="VUID-VkCopyAccelerationStructureInfoKHR-pNext-pNext"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-src-parameter"
  id="VUID-VkCopyAccelerationStructureInfoKHR-src-parameter"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-src-parameter  
  `src` **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-dst-parameter"
  id="VUID-VkCopyAccelerationStructureInfoKHR-dst-parameter"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-dst-parameter  
  `dst` **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-mode-parameter"
  id="VUID-VkCopyAccelerationStructureInfoKHR-mode-parameter"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-mode-parameter  
  `mode` **must** be a valid
  [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html)
  value

- <a href="#VUID-VkCopyAccelerationStructureInfoKHR-commonparent"
  id="VUID-VkCopyAccelerationStructureInfoKHR-commonparent"></a>
  VUID-VkCopyAccelerationStructureInfoKHR-commonparent  
  Both of `dst`, and `src` **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureModeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyAccelerationStructureKHR.html),
[vkCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyAccelerationStructureKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyAccelerationStructureInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
