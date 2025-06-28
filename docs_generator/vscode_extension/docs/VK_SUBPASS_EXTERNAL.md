# VK\_SUBPASS\_EXTERNAL(3) Manual Page

## Name

VK\_SUBPASS\_EXTERNAL - Subpass index sentinel expanding synchronization scope outside a subpass



## [](#_c_specification)C Specification

`VK_SUBPASS_EXTERNAL` is a special subpass index value expanding synchronization scope outside a subpass. It is described in more detail by [VkSubpassDependency](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency.html).

```c++
#define VK_SUBPASS_EXTERNAL               (~0U)
```

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_SUBPASS_EXTERNAL)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0