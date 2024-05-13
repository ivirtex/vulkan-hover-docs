# VkFenceCreateInfo(3) Manual Page

## Name

VkFenceCreateInfo - Structure specifying parameters of a newly created
fence



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFenceCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkFenceCreateInfo {
    VkStructureType       sType;
    const void*           pNext;
    VkFenceCreateFlags    flags;
} VkFenceCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkFenceCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateFlagBits.html) specifying the
  initial state and behavior of the fence.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkFenceCreateInfo-sType-sType"
  id="VUID-VkFenceCreateInfo-sType-sType"></a>
  VUID-VkFenceCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FENCE_CREATE_INFO`

- <a href="#VUID-VkFenceCreateInfo-pNext-pNext"
  id="VUID-VkFenceCreateInfo-pNext-pNext"></a>
  VUID-VkFenceCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of [VkExportFenceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceCreateInfo.html) or
  [VkExportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceWin32HandleInfoKHR.html)

- <a href="#VUID-VkFenceCreateInfo-sType-unique"
  id="VUID-VkFenceCreateInfo-sType-unique"></a>
  VUID-VkFenceCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkFenceCreateInfo-flags-parameter"
  id="VUID-VkFenceCreateInfo-flags-parameter"></a>
  VUID-VkFenceCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkFenceCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkFenceCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateFence.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFenceCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
