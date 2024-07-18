# VkBaseInStructure(3) Manual Page

## Name

VkBaseInStructure - Base structure for a read-only pointer chain



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBaseInStructure` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkBaseInStructure {
    VkStructureType                    sType;
    const struct VkBaseInStructure*    pNext;
} VkBaseInStructure;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is the structure type of the structure being iterated through.

- `pNext` is `NULL` or a pointer to the next structure in a structure
  chain.

## <a href="#_description" class="anchor"></a>Description

`VkBaseInStructure` can be used to facilitate iterating through a
read-only structure pointer chain.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBaseInStructure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBaseInStructure.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBaseInStructure"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
