# VkPipelineDiscardRectangleStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineDiscardRectangleStateCreateInfoEXT - Structure specifying discard rectangle



## [](#_c_specification)C Specification

The `VkPipelineDiscardRectangleStateCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_discard_rectangles
typedef struct VkPipelineDiscardRectangleStateCreateInfoEXT {
    VkStructureType                                  sType;
    const void*                                      pNext;
    VkPipelineDiscardRectangleStateCreateFlagsEXT    flags;
    VkDiscardRectangleModeEXT                        discardRectangleMode;
    uint32_t                                         discardRectangleCount;
    const VkRect2D*                                  pDiscardRectangles;
} VkPipelineDiscardRectangleStateCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `discardRectangleMode` is a [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDiscardRectangleModeEXT.html) value determining whether the discard rectangle test is inclusive or exclusive.
- `discardRectangleCount` is the number of discard rectangles to use.
- `pDiscardRectangles` is a pointer to an array of [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures defining discard rectangles.

## [](#_description)Description

If the `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` dynamic state is enabled for a pipeline, the `pDiscardRectangles` member is ignored. If the `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT` dynamic state is not enabled for the pipeline the presence of this structure in the [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) chain, and a `discardRectangleCount` greater than zero, implicitly enables discard rectangles in the pipeline, otherwise discard rectangles **must** enabled or disabled by [vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDiscardRectangleEnableEXT.html). If the `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT` dynamic state is enabled for the pipeline, the `discardRectangleMode` member is ignored, and the discard rectangle mode **must** be set by [vkCmdSetDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDiscardRectangleModeEXT.html).

When this structure is included in the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), it defines parameters of the discard rectangle test. If the `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` dynamic state is not enabled, and this structure is not included in the `pNext` chain, it is equivalent to specifying this structure with a `discardRectangleCount` of `0`. If all `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT`, `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT`, and `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT` dynamic states are enabled, the application **can** omit this structure from the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) and still use discard rectangles by setting all state dynamically. In this case `vkCmdSetDiscardRectangleEXT` **must** be called to set the discard rectangle for all indices \[0, `maxDiscardRectangles`) before drawing with discard rectangles enabled. Individual discard rectangles **can** be made ineffective by setting their [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html)::`extent.width` and [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html)::`extent.height` to zero.

Valid Usage

- [](#VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleCount-00582)VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleCount-00582  
  `discardRectangleCount` **must** be less than or equal to `VkPhysicalDeviceDiscardRectanglePropertiesEXT`::`maxDiscardRectangles`

Valid Usage (Implicit)

- [](#VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-sType-sType)VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT`
- [](#VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-flags-zerobitmask)VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleMode-parameter)VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleMode-parameter  
  `discardRectangleMode` **must** be a valid [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDiscardRectangleModeEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_discard\_rectangles](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_discard_rectangles.html), [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDiscardRectangleModeEXT.html), [VkPipelineDiscardRectangleStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDiscardRectangleStateCreateFlagsEXT.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineDiscardRectangleStateCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0