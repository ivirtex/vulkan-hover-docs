# VkLatencySubmissionPresentIdNV(3) Manual Page

## Name

VkLatencySubmissionPresentIdNV - Structure used to associate a
queueSubmit with a presentId



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySubmissionPresentIdNV.html)
structure is defined as:

``` c
// Provided by VK_NV_low_latency2
typedef struct VkLatencySubmissionPresentIdNV {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           presentID;
} VkLatencySubmissionPresentIdNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `presentId` is used to associate the `vkQueueSubmit` with the
  presentId used for a given `vkQueuePresentKHR` via
  [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentIdKHR.html)::`pPresentIds`.

## <a href="#_description" class="anchor"></a>Description

For any submission to be tracked with low latency mode pacing, it needs
to be associated with other submissions in a given present. Applications
**must** include the VkLatencySubmissionPresentIdNV in the pNext chain
of [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html) to associate that submission with
the `presentId` present for low latency mode.

Valid Usage (Implicit)

- <a href="#VUID-VkLatencySubmissionPresentIdNV-sType-sType"
  id="VUID-VkLatencySubmissionPresentIdNV-sType-sType"></a>
  VUID-VkLatencySubmissionPresentIdNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_LATENCY_SUBMISSION_PRESENT_ID_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLatencySubmissionPresentIdNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
