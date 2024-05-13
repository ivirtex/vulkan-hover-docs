# VkCullModeFlagBits(3) Manual Page

## Name

VkCullModeFlagBits - Bitmask controlling triangle culling



## <a href="#_c_specification" class="anchor"></a>C Specification

Once the orientation of triangles is determined, they are culled
according to the
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)::`cullMode`
property of the currently active pipeline. Possible values are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkCullModeFlagBits {
    VK_CULL_MODE_NONE = 0,
    VK_CULL_MODE_FRONT_BIT = 0x00000001,
    VK_CULL_MODE_BACK_BIT = 0x00000002,
    VK_CULL_MODE_FRONT_AND_BACK = 0x00000003,
} VkCullModeFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_CULL_MODE_NONE` specifies that no triangles are discarded

- `VK_CULL_MODE_FRONT_BIT` specifies that front-facing triangles are
  discarded

- `VK_CULL_MODE_BACK_BIT` specifies that back-facing triangles are
  discarded

- `VK_CULL_MODE_FRONT_AND_BACK` specifies that all triangles are
  discarded.

Following culling, fragments are produced for any triangles which have
not been discarded.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCullModeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCullModeFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCullModeFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
