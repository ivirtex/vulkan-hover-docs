# VkTileShadingRenderPassFlagBitsQCOM(3) Manual Page

## Name

VkTileShadingRenderPassFlagBitsQCOM - Bitmask specifying flags for tile shading



## [](#_c_specification)C Specification

Bits which **can** be set in [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`flags` describing additional properties of the render pass are:

```c++
// Provided by VK_QCOM_tile_shading
typedef enum VkTileShadingRenderPassFlagBitsQCOM {
    VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM = 0x00000001,
    VK_TILE_SHADING_RENDER_PASS_PER_TILE_EXECUTION_BIT_QCOM = 0x00000002,
} VkTileShadingRenderPassFlagBitsQCOM;
```

## [](#_description)Description

- `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` specifies that the render pass has tile shading enabled.
- `VK_TILE_SHADING_RENDER_PASS_PER_TILE_EXECUTION_BIT_QCOM` specifies that the secondary command buffer will be executed within a [per-tile execution block](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model).

## [](#_see_also)See Also

[VK\_QCOM\_tile\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_shading.html), [VkTileShadingRenderPassFlagsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileShadingRenderPassFlagsQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTileShadingRenderPassFlagBitsQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0