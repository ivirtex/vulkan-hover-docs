# VkAmigoProfilingSubmitInfoSEC(3) Manual Page

## Name

VkAmigoProfilingSubmitInfoSEC - Stub description of VkAmigoProfilingSubmitInfoSEC



## [](#_c_specification)C Specification

There is currently no specification language written for this type. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_SEC_amigo_profiling
typedef struct VkAmigoProfilingSubmitInfoSEC {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           firstDrawTimestamp;
    uint64_t           swapBufferTimestamp;
} VkAmigoProfilingSubmitInfoSEC;
```

## [](#_members)Members

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkAmigoProfilingSubmitInfoSEC-sType-sType)VUID-VkAmigoProfilingSubmitInfoSEC-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_AMIGO_PROFILING_SUBMIT_INFO_SEC`

## [](#_see_also)See Also

[VK\_SEC\_amigo\_profiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_SEC_amigo_profiling.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAmigoProfilingSubmitInfoSEC)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0