# VkPipelineDiscardRectangleStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineDiscardRectangleStateCreateInfoEXT - Structure specifying
discard rectangle



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineDiscardRectangleStateCreateInfoEXT` structure is defined
as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `discardRectangleMode` is a
  [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDiscardRectangleModeEXT.html) value
  determining whether the discard rectangle test is inclusive or
  exclusive.

- `discardRectangleCount` is the number of discard rectangles to use.

- `pDiscardRectangles` is a pointer to an array of
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures defining discard rectangles.

## <a href="#_description" class="anchor"></a>Description

If the `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` dynamic state is enabled
for a pipeline, the `pDiscardRectangles` member is ignored. If the
`VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT` dynamic state is not
enabled for the pipeline the presence of this structure in the
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html) chain,
and a `discardRectangleCount` greater than zero, implicitly enables
discard rectangles in the pipeline, otherwise discard rectangles
**must** enabled or disabled by
[vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEnableEXT.html).
If the `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT` dynamic state is
enabled for the pipeline, the `discardRectangleMode` member is ignored,
and the discard rectangle mode **must** be set by
[vkCmdSetDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleModeEXT.html).

When this structure is included in the `pNext` chain of
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html), it
defines parameters of the discard rectangle test. If the
`VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` dynamic state is not enabled,
and this structure is not included in the `pNext` chain, it is
equivalent to specifying this structure with a `discardRectangleCount`
of `0`.

Valid Usage

- <a
  href="#VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleCount-00582"
  id="VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleCount-00582"></a>
  VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleCount-00582  
  `discardRectangleCount` **must** be less than or equal to
  `VkPhysicalDeviceDiscardRectanglePropertiesEXT`::`maxDiscardRectangles`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_DISCARD_RECTANGLE_STATE_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleMode-parameter"
  id="VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleMode-parameter"></a>
  VUID-VkPipelineDiscardRectangleStateCreateInfoEXT-discardRectangleMode-parameter  
  `discardRectangleMode` **must** be a valid
  [VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDiscardRectangleModeEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_discard_rectangles](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_discard_rectangles.html),
[VkDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDiscardRectangleModeEXT.html),
[VkPipelineDiscardRectangleStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateFlagsEXT.html),
[VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineDiscardRectangleStateCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
