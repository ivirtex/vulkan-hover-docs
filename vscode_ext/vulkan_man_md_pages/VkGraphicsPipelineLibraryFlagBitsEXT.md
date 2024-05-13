# VkGraphicsPipelineLibraryFlagBitsEXT(3) Manual Page

## Name

VkGraphicsPipelineLibraryFlagBitsEXT - Bitmask specifying the subset of
a graphics pipeline to compile



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the `flags` member of
[VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html),
specifying the subsets of a graphics pipeline to compile are:

``` c
// Provided by VK_EXT_graphics_pipeline_library
typedef enum VkGraphicsPipelineLibraryFlagBitsEXT {
    VK_GRAPHICS_PIPELINE_LIBRARY_VERTEX_INPUT_INTERFACE_BIT_EXT = 0x00000001,
    VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT = 0x00000002,
    VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT = 0x00000004,
    VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT = 0x00000008,
} VkGraphicsPipelineLibraryFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_GRAPHICS_PIPELINE_LIBRARY_VERTEX_INPUT_INTERFACE_BIT_EXT`
  specifies that a pipeline will include <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input interface state</a>.

- `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`
  specifies that a pipeline will include <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>.

- `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT` specifies that
  a pipeline will include <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>.

- `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`
  specifies that a pipeline will include <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_graphics_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_graphics_pipeline_library.html),
[VkGraphicsPipelineLibraryFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGraphicsPipelineLibraryFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
