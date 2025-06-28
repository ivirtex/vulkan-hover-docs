# VkAntiLagDataAMD(3) Manual Page

## Name

VkAntiLagDataAMD - Structure specifying the parameters for vkAntiLagUpdateAMD



## [](#_c_specification)C Specification

The `VkAntiLagDataAMD` structure is defined as:

```c++
// Provided by VK_AMD_anti_lag
typedef struct VkAntiLagDataAMD {
    VkStructureType                        sType;
    const void*                            pNext;
    VkAntiLagModeAMD                       mode;
    uint32_t                               maxFPS;
    const VkAntiLagPresentationInfoAMD*    pPresentationInfo;
} VkAntiLagDataAMD;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `mode` is a [VkAntiLagModeAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagModeAMD.html) value specifying the anti-lag status.
- `maxFPS` is the framerate limit, in frames per second, used by the application. This limit will be imposed if anti-lag is enabled. If the application tries to render faster, the framerate will be reduced to match this limit. A value of 0 will disable the limit.
- `pPresentationInfo` is a pointer to a [VkAntiLagPresentationInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagPresentationInfoAMD.html) structure containing information about the application stage.

## [](#_description)Description

This structure specifies anti-lag parameters.

Valid Usage (Implicit)

- [](#VUID-VkAntiLagDataAMD-sType-sType)VUID-VkAntiLagDataAMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ANTI_LAG_DATA_AMD`
- [](#VUID-VkAntiLagDataAMD-mode-parameter)VUID-VkAntiLagDataAMD-mode-parameter  
  `mode` **must** be a valid [VkAntiLagModeAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagModeAMD.html) value
- [](#VUID-VkAntiLagDataAMD-pPresentationInfo-parameter)VUID-VkAntiLagDataAMD-pPresentationInfo-parameter  
  If `pPresentationInfo` is not `NULL`, `pPresentationInfo` **must** be a valid pointer to a valid [VkAntiLagPresentationInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagPresentationInfoAMD.html) structure

## [](#_see_also)See Also

[VK\_AMD\_anti\_lag](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_anti_lag.html), [VkAntiLagModeAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagModeAMD.html), [VkAntiLagPresentationInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagPresentationInfoAMD.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkAntiLagUpdateAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAntiLagUpdateAMD.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAntiLagDataAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0