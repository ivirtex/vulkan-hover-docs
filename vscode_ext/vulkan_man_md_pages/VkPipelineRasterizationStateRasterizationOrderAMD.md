# VkPipelineRasterizationStateRasterizationOrderAMD(3) Manual Page

## Name

VkPipelineRasterizationStateRasterizationOrderAMD - Structure defining
rasterization order for a graphics pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The rasterization order to use for a graphics pipeline is specified by
adding a `VkPipelineRasterizationStateRasterizationOrderAMD` structure
to the `pNext` chain of a
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
structure.

The `VkPipelineRasterizationStateRasterizationOrderAMD` structure is
defined as:

``` c
// Provided by VK_AMD_rasterization_order
typedef struct VkPipelineRasterizationStateRasterizationOrderAMD {
    VkStructureType            sType;
    const void*                pNext;
    VkRasterizationOrderAMD    rasterizationOrder;
} VkPipelineRasterizationStateRasterizationOrderAMD;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `rasterizationOrder` is a
  [VkRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRasterizationOrderAMD.html) value
  specifying the primitive rasterization order to use.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineRasterizationStateRasterizationOrderAMD-sType-sType"
  id="VUID-VkPipelineRasterizationStateRasterizationOrderAMD-sType-sType"></a>
  VUID-VkPipelineRasterizationStateRasterizationOrderAMD-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD`

- <a
  href="#VUID-VkPipelineRasterizationStateRasterizationOrderAMD-rasterizationOrder-parameter"
  id="VUID-VkPipelineRasterizationStateRasterizationOrderAMD-rasterizationOrder-parameter"></a>
  VUID-VkPipelineRasterizationStateRasterizationOrderAMD-rasterizationOrder-parameter  
  `rasterizationOrder` **must** be a valid
  [VkRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRasterizationOrderAMD.html) value

If the [`VK_AMD_rasterization_order`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_rasterization_order.html)
device extension is not enabled or the application does not request a
particular rasterization order through specifying a
`VkPipelineRasterizationStateRasterizationOrderAMD` structure then the
rasterization order used by the graphics pipeline defaults to
`VK_RASTERIZATION_ORDER_STRICT_AMD`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_rasterization_order](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_rasterization_order.html),
[VkRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRasterizationOrderAMD.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRasterizationStateRasterizationOrderAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
