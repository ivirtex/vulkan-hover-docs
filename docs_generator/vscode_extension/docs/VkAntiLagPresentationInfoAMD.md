# VkAntiLagPresentationInfoAMD(3) Manual Page

## Name

VkAntiLagPresentationInfoAMD - Structure specifying information about stage



## [](#_c_specification)C Specification

The `VkAntiLagPresentationInfoAMD` structure is defined as:

```c++
// Provided by VK_AMD_anti_lag
typedef struct VkAntiLagPresentationInfoAMD {
    VkStructureType      sType;
    void*                pNext;
    VkAntiLagStageAMD    stage;
    uint64_t             frameIndex;
} VkAntiLagPresentationInfoAMD;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stage` is a [VkAntiLagStageAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagStageAMD.html) value specifying the current application stage.
- `frameIndex` is set just before the application processes input data (`VK_ANTI_LAG_STAGE_INPUT_AMD`). The same `frameIndex` value **should** be set before the frame with current input data will be presented by [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html) (`VK_ANTI_LAG_STAGE_PRESENT_AMD`). This **should** be done for each frame.

## [](#_description)Description

This structure specifies information about the presentation stage for which anti-lag parameters are being set.

Valid Usage (Implicit)

- [](#VUID-VkAntiLagPresentationInfoAMD-sType-sType)VUID-VkAntiLagPresentationInfoAMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ANTI_LAG_PRESENTATION_INFO_AMD`
- [](#VUID-VkAntiLagPresentationInfoAMD-stage-parameter)VUID-VkAntiLagPresentationInfoAMD-stage-parameter  
  `stage` **must** be a valid [VkAntiLagStageAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagStageAMD.html) value

## [](#_see_also)See Also

[VK\_AMD\_anti\_lag](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_anti_lag.html), [VkAntiLagDataAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagDataAMD.html), [VkAntiLagStageAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagStageAMD.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAntiLagPresentationInfoAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0