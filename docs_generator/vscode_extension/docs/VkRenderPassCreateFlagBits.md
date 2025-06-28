# VkRenderPassCreateFlagBits(3) Manual Page

## Name

VkRenderPassCreateFlagBits - Bitmask specifying additional properties of a render pass



## [](#_c_specification)C Specification

Bits which **can** be set in [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html)::`flags`, describing additional properties of the render pass, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkRenderPassCreateFlagBits {
  // Provided by VK_QCOM_render_pass_transform
    VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM = 0x00000002,
  // Provided by VK_VALVE_fragment_density_map_layered
    VK_RENDER_PASS_CREATE_PER_LAYER_FRAGMENT_DENSITY_BIT_VALVE = 0x00000004,
} VkRenderPassCreateFlagBits;
```

## [](#_description)Description

- `VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM` specifies that the created render pass is compatible with [render pass transform](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-renderpass-transform).
- `VK_RENDER_PASS_CREATE_PER_LAYER_FRAGMENT_DENSITY_BIT_VALVE` specifies that the created render pass is usable with layered fragment density maps.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkRenderPassCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassCreateFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0