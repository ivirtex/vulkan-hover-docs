# VK_REMAINING_ARRAY_LAYERS(3) Manual Page

## Name

VK_REMAINING_ARRAY_LAYERS - Sentinel for all remaining array layers



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_REMAINING_ARRAY_LAYERS` is a special constant value used for image
views to indicate that all remaining array layers in an image after the
base layer should be included in the view.

``` c
#define VK_REMAINING_ARRAY_LAYERS         (~0U)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_REMAINING_ARRAY_LAYERS"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
