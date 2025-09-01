# VkLatencySubmissionPresentIdNV(3) Manual Page

## Name

VkLatencySubmissionPresentIdNV - Structure used to associate a queueSubmit with a presentId



## [](#_c_specification)C Specification

The [VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySubmissionPresentIdNV.html) structure is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef struct VkLatencySubmissionPresentIdNV {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           presentID;
} VkLatencySubmissionPresentIdNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentID` is used to associate the `vkQueueSubmit` with the presentId used for a given `vkQueuePresentKHR` via [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentIdKHR.html)::`pPresentIds` or [VkPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentId2KHR.html)::`pPresentIds`.

## [](#_description)Description

For any submission to be tracked with low latency mode pacing, it needs to be associated with other submissions in a given present. To associate a submission with `presentID` for low latency mode, the `pNext` chain of [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit.html) **must** include a `VkLatencySubmissionPresentIdNV` structure.

Valid Usage (Implicit)

- [](#VUID-VkLatencySubmissionPresentIdNV-sType-sType)VUID-VkLatencySubmissionPresentIdNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_LATENCY_SUBMISSION_PRESENT_ID_NV`

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLatencySubmissionPresentIdNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0