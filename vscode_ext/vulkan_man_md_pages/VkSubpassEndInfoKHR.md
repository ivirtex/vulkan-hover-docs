# VkSubpassEndInfo(3) Manual Page

## Name

VkSubpassEndInfo - Structure specifying subpass end information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSubpassEndInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkSubpassEndInfo {
    VkStructureType    sType;
    const void*        pNext;
} VkSubpassEndInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_create_renderpass2
typedef VkSubpassEndInfo VkSubpassEndInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkSubpassEndInfo-sType-sType"
  id="VUID-VkSubpassEndInfo-sType-sType"></a>
  VUID-VkSubpassEndInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_END_INFO`

- <a href="#VUID-VkSubpassEndInfo-pNext-pNext"
  id="VUID-VkSubpassEndInfo-pNext-pNext"></a>
  VUID-VkSubpassEndInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkSubpassFragmentDensityMapOffsetEndInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassFragmentDensityMapOffsetEndInfoQCOM.html)

- <a href="#VUID-VkSubpassEndInfo-sType-unique"
  id="VUID-VkSubpassEndInfo-sType-unique"></a>
  VUID-VkSubpassEndInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdEndRenderPass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass2.html),
[vkCmdEndRenderPass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass2KHR.html),
[vkCmdNextSubpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass2.html),
[vkCmdNextSubpass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassEndInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
