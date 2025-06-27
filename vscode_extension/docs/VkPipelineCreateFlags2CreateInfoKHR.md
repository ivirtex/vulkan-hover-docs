# VkPipelineCreateFlags2CreateInfoKHR(3) Manual Page

## Name

VkPipelineCreateFlags2CreateInfoKHR - Extended pipeline create flags



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineCreateFlags2CreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance5
typedef struct VkPipelineCreateFlags2CreateInfoKHR {
    VkStructureType              sType;
    const void*                  pNext;
    VkPipelineCreateFlags2KHR    flags;
} VkPipelineCreateFlags2CreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html)
  specifying how a pipeline will be generated.

## <a href="#_description" class="anchor"></a>Description

If this structure is included in the `pNext` chain of a pipeline
creation structure, `flags` is used instead of the corresponding `flags`
value passed in that creation structure, allowing additional creation
flags to be specified.

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineCreateFlags2CreateInfoKHR-sType-sType"
  id="VUID-VkPipelineCreateFlags2CreateInfoKHR-sType-sType"></a>
  VUID-VkPipelineCreateFlags2CreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_CREATE_FLAGS_2_CREATE_INFO_KHR`

- <a href="#VUID-VkPipelineCreateFlags2CreateInfoKHR-flags-parameter"
  id="VUID-VkPipelineCreateFlags2CreateInfoKHR-flags-parameter"></a>
  VUID-VkPipelineCreateFlags2CreateInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of
  [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html)
  values

- <a
  href="#VUID-VkPipelineCreateFlags2CreateInfoKHR-flags-requiredbitmask"
  id="VUID-VkPipelineCreateFlags2CreateInfoKHR-flags-requiredbitmask"></a>
  VUID-VkPipelineCreateFlags2CreateInfoKHR-flags-requiredbitmask  
  `flags` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkPipelineCreateFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2KHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCreateFlags2CreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
