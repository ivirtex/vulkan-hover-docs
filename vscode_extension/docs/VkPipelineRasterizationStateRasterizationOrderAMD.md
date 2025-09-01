# VkPipelineRasterizationStateRasterizationOrderAMD(3) Manual Page

## Name

VkPipelineRasterizationStateRasterizationOrderAMD - Structure defining rasterization order for a graphics pipeline



## [](#_c_specification)C Specification

The rasterization order to use for a graphics pipeline is specified by adding a `VkPipelineRasterizationStateRasterizationOrderAMD` structure to the `pNext` chain of a [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html) structure.

The `VkPipelineRasterizationStateRasterizationOrderAMD` structure is defined as:

```c++
// Provided by VK_AMD_rasterization_order
typedef struct VkPipelineRasterizationStateRasterizationOrderAMD {
    VkStructureType            sType;
    const void*                pNext;
    VkRasterizationOrderAMD    rasterizationOrder;
} VkPipelineRasterizationStateRasterizationOrderAMD;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `rasterizationOrder` is a [VkRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRasterizationOrderAMD.html) value specifying the primitive rasterization order to use.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineRasterizationStateRasterizationOrderAMD-sType-sType)VUID-VkPipelineRasterizationStateRasterizationOrderAMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_RASTERIZATION_ORDER_AMD`
- [](#VUID-VkPipelineRasterizationStateRasterizationOrderAMD-rasterizationOrder-parameter)VUID-VkPipelineRasterizationStateRasterizationOrderAMD-rasterizationOrder-parameter  
  `rasterizationOrder` **must** be a valid [VkRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRasterizationOrderAMD.html) value

If the `VK_AMD_rasterization_order` device extension is not enabled or the application does not request a particular rasterization order through specifying a `VkPipelineRasterizationStateRasterizationOrderAMD` structure then the rasterization order used by the graphics pipeline defaults to `VK_RASTERIZATION_ORDER_STRICT_AMD`.

## [](#_see_also)See Also

[VK\_AMD\_rasterization\_order](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_rasterization_order.html), [VkRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRasterizationOrderAMD.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRasterizationStateRasterizationOrderAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0