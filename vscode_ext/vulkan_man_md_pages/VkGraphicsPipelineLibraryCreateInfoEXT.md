# VkGraphicsPipelineLibraryCreateInfoEXT(3) Manual Page

## Name

VkGraphicsPipelineLibraryCreateInfoEXT - Structure specifying the
subsets of the graphics pipeline being compiled



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGraphicsPipelineLibraryCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_graphics_pipeline_library
typedef struct VkGraphicsPipelineLibraryCreateInfoEXT {
    VkStructureType                      sType;
    const void*                          pNext;
    VkGraphicsPipelineLibraryFlagsEXT    flags;
} VkGraphicsPipelineLibraryCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkGraphicsPipelineLibraryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryFlagBitsEXT.html)
  specifying the subsets of the graphics pipeline that are being
  compiled.

## <a href="#_description" class="anchor"></a>Description

If a `VkGraphicsPipelineLibraryCreateInfoEXT` structure is included in
the `pNext` chain of
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html), it
specifies the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets"
target="_blank" rel="noopener">subsets of the graphics pipeline</a>
being created, excluding any subsets from linked pipeline libraries. If
the pipeline is created with pipeline libraries, state from those
libraries is aggregated with said subset.

If this structure is omitted, and either
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`flags`
includes `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` or the
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pNext`
chain includes a
[VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
structure with a `libraryCount` greater than `0`, it is as if `flags` is
`0`. Otherwise if this structure is omitted, it is as if `flags`
includes all possible subsets of the graphics pipeline (i.e. a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
target="_blank" rel="noopener">complete graphics pipeline</a>).

Valid Usage (Implicit)

- <a href="#VUID-VkGraphicsPipelineLibraryCreateInfoEXT-sType-sType"
  id="VUID-VkGraphicsPipelineLibraryCreateInfoEXT-sType-sType"></a>
  VUID-VkGraphicsPipelineLibraryCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_LIBRARY_CREATE_INFO_EXT`

- <a href="#VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-parameter"
  id="VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-parameter"></a>
  VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkGraphicsPipelineLibraryFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryFlagBitsEXT.html)
  values

- <a
  href="#VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-requiredbitmask"
  id="VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-requiredbitmask"></a>
  VUID-VkGraphicsPipelineLibraryCreateInfoEXT-flags-requiredbitmask  
  `flags` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_graphics_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_graphics_pipeline_library.html),
[VkGraphicsPipelineLibraryFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGraphicsPipelineLibraryCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
