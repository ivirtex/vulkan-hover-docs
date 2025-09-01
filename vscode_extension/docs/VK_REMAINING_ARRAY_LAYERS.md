# VK\_REMAINING\_ARRAY\_LAYERS(3) Manual Page

## Name

VK\_REMAINING\_ARRAY\_LAYERS - Sentinel for all remaining array layers



## [](#_c_specification)C Specification

`VK_REMAINING_ARRAY_LAYERS` is a special constant value used for image views to indicate that all remaining array layers in an image after the base layer should be included in the view.

```c++
#define VK_REMAINING_ARRAY_LAYERS         (~0U)
```

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_REMAINING_ARRAY_LAYERS).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0