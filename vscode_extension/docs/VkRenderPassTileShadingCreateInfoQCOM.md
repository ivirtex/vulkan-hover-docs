# VkRenderPassTileShadingCreateInfoQCOM(3) Manual Page

## Name

VkRenderPassTileShadingCreateInfoQCOM - Structure specifying, tile shading information for a render pass object.



## [](#_c_specification)C Specification

To enable tile shading for a render pass object, add a [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html) to the `pNext` chain of [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html) or [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html) . To enable tile shading for a dynamic render pass, add a [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html) to the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html). To execute a secondary command buffer within a render pass, add a [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html) to the `pNext` chain of [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html) when the secondary command buffer is recorded.

The `VkRenderPassTileShadingCreateInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_shading
typedef struct VkRenderPassTileShadingCreateInfoQCOM {
    VkStructureType                     sType;
    const void*                         pNext;
    VkTileShadingRenderPassFlagsQCOM    flags;
    VkExtent2D                          tileApronSize;
} VkRenderPassTileShadingCreateInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkTileShadingRenderPassFlagBitsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileShadingRenderPassFlagBitsQCOM.html).
- `tileApronSize` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) describing the is size of the [tiling apron](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading-aprons) in each dimension.

## [](#_description)Description

If this structure is not present, the render pass will have `flags` set to `0` and `tileApronSize` is set to `(0,0)`.

Valid Usage

- [](#VUID-VkRenderPassTileShadingCreateInfoQCOM-tileShading-10658)VUID-VkRenderPassTileShadingCreateInfoQCOM-tileShading-10658  
  If the [`tileShading`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileShading) feature is not enabled, `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` **must** not be included in `flags`
- [](#VUID-VkRenderPassTileShadingCreateInfoQCOM-flags-10659)VUID-VkRenderPassTileShadingCreateInfoQCOM-flags-10659  
  If `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` is not included in `flags` or the [`tileShadingApron`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileShadingApron) feature is not enabled, `tileApronSize` **must** be `(0,0)`
- [](#VUID-VkRenderPassTileShadingCreateInfoQCOM-flags-10660)VUID-VkRenderPassTileShadingCreateInfoQCOM-flags-10660  
  If `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` is not included in `flags`, or neither the [tileShadingPerTileDispatch](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileShadingPerTileDispatch) and [tileShadingPerTileDraw](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileShadingPerTileDraw) features are enabled, `flags` **must** not include `VK_TILE_SHADING_RENDER_PASS_PER_TILE_EXECUTION_BIT_QCOM`
- [](#VUID-VkRenderPassTileShadingCreateInfoQCOM-tileShadingAnisotropicApron-10661)VUID-VkRenderPassTileShadingCreateInfoQCOM-tileShadingAnisotropicApron-10661  
  If the [`tileShadingAnisotropicApron`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileShadingAnisotropicApron) feature is not enabled, `tileApronSize.x` and **must** be equal to `tileApronSize.y`
- [](#VUID-VkRenderPassTileShadingCreateInfoQCOM-tileApronSize-10662)VUID-VkRenderPassTileShadingCreateInfoQCOM-tileApronSize-10662  
  `tileApronSize.x` **must** be less than or equal to [`maxApronSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxApronSize)
- [](#VUID-VkRenderPassTileShadingCreateInfoQCOM-tileApronSize-10663)VUID-VkRenderPassTileShadingCreateInfoQCOM-tileApronSize-10663  
  `tileApronSize.y` **must** be less than or equal to [`maxApronSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxApronSize)

Valid Usage (Implicit)

- [](#VUID-VkRenderPassTileShadingCreateInfoQCOM-sType-sType)VUID-VkRenderPassTileShadingCreateInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_TILE_SHADING_CREATE_INFO_QCOM`
- [](#VUID-VkRenderPassTileShadingCreateInfoQCOM-flags-parameter)VUID-VkRenderPassTileShadingCreateInfoQCOM-flags-parameter  
  `flags` **must** be a valid combination of [VkTileShadingRenderPassFlagBitsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileShadingRenderPassFlagBitsQCOM.html) values

## [](#_see_also)See Also

[VK\_QCOM\_tile\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_shading.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTileShadingRenderPassFlagsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileShadingRenderPassFlagsQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassTileShadingCreateInfoQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0