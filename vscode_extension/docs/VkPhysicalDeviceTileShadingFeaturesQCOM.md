# VkPhysicalDeviceTileShadingFeaturesQCOM(3) Manual Page

## Name

VkPhysicalDeviceTileShadingFeaturesQCOM - Structure describing tile shading features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceTileShadingFeaturesQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_shading
typedef struct VkPhysicalDeviceTileShadingFeaturesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           tileShading;
    VkBool32           tileShadingFragmentStage;
    VkBool32           tileShadingColorAttachments;
    VkBool32           tileShadingDepthAttachments;
    VkBool32           tileShadingStencilAttachments;
    VkBool32           tileShadingInputAttachments;
    VkBool32           tileShadingSampledAttachments;
    VkBool32           tileShadingPerTileDraw;
    VkBool32           tileShadingPerTileDispatch;
    VkBool32           tileShadingDispatchTile;
    VkBool32           tileShadingApron;
    VkBool32           tileShadingAnisotropicApron;
    VkBool32           tileShadingAtomicOps;
    VkBool32           tileShadingImageProcessing;
} VkPhysicalDeviceTileShadingFeaturesQCOM;
```

## [](#_members)Members

This structure describes the following features:

- []()`tileShading` indicates that the implementation supports [tile shading render pass](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading) instances.
- []()`tileShadingFragmentStage` indicates that the implementation supports tile shading in the fragment stage.
- []()`tileShadingColorAttachments` indicates that the implementation supports access to color attachments in a tile shader.
- []()`tileShadingDepthAttachments` indicates that the implementation supports access to depth aspect of depth stencil attachments.
- []()`tileShadingStencilAttachments` indicates that the implementation supports access to stencil aspect of depth stencil attachments.
- []()`tileShadingInputAttachments` indicates that the implementation supports access to input attachments.
- []()`tileShadingSampledAttachments` indicates that the implementation supports access to sampling of tile attachments.
- []()`tileShadingPerTileDraw` indicates that the implementation supports the recording of vkCmdDraw* commands when [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled.
- []()`tileShadingPerTileDispatch` indicates that the implementation supports the recording of `vkCmdDispatch`* commands within those regions of a command buffer where the [per-tile execution model](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-per-tile-execution-model) is enabled.
- []()`tileShadingDispatchTile` indicates that the implementation supports the recording of [vkCmdDispatchTileQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchTileQCOM.html) commands.
- []()`tileShadingApron` indicates that the implementation supports [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`apronSize` value other than (0,0). See [Tiling Aprons](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-tile-shading-aprons) for more information.
- []()`tileShadingAnisotropicApron` indicates that the implementation supports [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`apronSize` set to a value where `apronSize.width` differs from `apronSize.height`.
- []()`tileShadingAtomicOps` indicates that the implementation supports atomic operations on [tile attachment variables](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-tile-attachment).
- []()`tileShadingImageProcessing` indicates that the implementation supports [image processing operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-weightimage) with tile attachments.

## [](#_description)Description

If the `VkPhysicalDeviceTileShadingFeaturesQCOM` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceTileShadingFeaturesQCOM`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceTileShadingFeaturesQCOM-sType-sType)VUID-VkPhysicalDeviceTileShadingFeaturesQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TILE_SHADING_FEATURES_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_tile\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_shading.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceTileShadingFeaturesQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0