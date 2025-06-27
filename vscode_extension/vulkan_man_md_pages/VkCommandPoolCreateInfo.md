# VkCommandPoolCreateInfo(3) Manual Page

## Name

VkCommandPoolCreateInfo - Structure specifying parameters of a newly
created command pool



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCommandPoolCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkCommandPoolCreateInfo {
    VkStructureType             sType;
    const void*                 pNext;
    VkCommandPoolCreateFlags    flags;
    uint32_t                    queueFamilyIndex;
} VkCommandPoolCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkCommandPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateFlagBits.html)
  indicating usage behavior for the pool and command buffers allocated
  from it.

- `queueFamilyIndex` designates a queue family as described in section
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-queueprops"
  target="_blank" rel="noopener">Queue Family Properties</a>. All
  command buffers allocated from this command pool **must** be submitted
  on queues from the same queue family.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCommandPoolCreateInfo-flags-02860"
  id="VUID-VkCommandPoolCreateInfo-flags-02860"></a>
  VUID-VkCommandPoolCreateInfo-flags-02860  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-protectedMemory"
  target="_blank" rel="noopener"><code>protectedMemory</code></a>
  feature is not enabled, the `VK_COMMAND_POOL_CREATE_PROTECTED_BIT` bit
  of `flags` **must** not be set

Valid Usage (Implicit)

- <a href="#VUID-VkCommandPoolCreateInfo-sType-sType"
  id="VUID-VkCommandPoolCreateInfo-sType-sType"></a>
  VUID-VkCommandPoolCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO`

- <a href="#VUID-VkCommandPoolCreateInfo-pNext-pNext"
  id="VUID-VkCommandPoolCreateInfo-pNext-pNext"></a>
  VUID-VkCommandPoolCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCommandPoolCreateInfo-flags-parameter"
  id="VUID-VkCommandPoolCreateInfo-flags-parameter"></a>
  VUID-VkCommandPoolCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkCommandPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandPoolCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCommandPool.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandPoolCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
