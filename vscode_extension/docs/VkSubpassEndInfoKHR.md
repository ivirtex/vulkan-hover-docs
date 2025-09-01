# VkSubpassEndInfo(3) Manual Page

## Name

VkSubpassEndInfo - Structure specifying subpass end information



## [](#_c_specification)C Specification

The `VkSubpassEndInfo` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
typedef struct VkSubpassEndInfo {
    VkStructureType    sType;
    const void*        pNext;
} VkSubpassEndInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_create_renderpass2
typedef VkSubpassEndInfo VkSubpassEndInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSubpassEndInfo-sType-sType)VUID-VkSubpassEndInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_END_INFO`
- [](#VUID-VkSubpassEndInfo-pNext-pNext)VUID-VkSubpassEndInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkRenderPassFragmentDensityMapOffsetEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapOffsetEndInfoEXT.html)
- [](#VUID-VkSubpassEndInfo-sType-unique)VUID-VkSubpassEndInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdEndRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRenderPass2.html), [vkCmdEndRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRenderPass2.html), [vkCmdNextSubpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdNextSubpass2.html), [vkCmdNextSubpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdNextSubpass2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassEndInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0