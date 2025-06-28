# VkSubpassBeginInfo(3) Manual Page

## Name

VkSubpassBeginInfo - Structure specifying subpass begin information



## [](#_c_specification)C Specification

The `VkSubpassBeginInfo` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
typedef struct VkSubpassBeginInfo {
    VkStructureType      sType;
    const void*          pNext;
    VkSubpassContents    contents;
} VkSubpassBeginInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_create_renderpass2
typedef VkSubpassBeginInfo VkSubpassBeginInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `contents` is a [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html) value specifying how the commands in the next subpass will be provided.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSubpassBeginInfo-contents-09382)VUID-VkSubpassBeginInfo-contents-09382  
  If `contents` is `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_KHR`, then at least one of the following features **must** be enabled:
  
  - [`maintenance7`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance7)
  - [`nestedCommandBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nestedCommandBuffer)

Valid Usage (Implicit)

- [](#VUID-VkSubpassBeginInfo-sType-sType)VUID-VkSubpassBeginInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO`
- [](#VUID-VkSubpassBeginInfo-pNext-pNext)VUID-VkSubpassBeginInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkSubpassBeginInfo-contents-parameter)VUID-VkSubpassBeginInfo-contents-parameter  
  `contents` **must** be a valid [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html) value

## [](#_see_also)See Also

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html), [vkCmdBeginRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass2.html), [vkCmdBeginRenderPass2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass2KHR.html), [vkCmdNextSubpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdNextSubpass2.html), [vkCmdNextSubpass2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdNextSubpass2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassBeginInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0