# VkRenderPass(3) Manual Page

## Name

VkRenderPass - Opaque handle to a render pass object



## <a href="#_c_specification" class="anchor"></a>C Specification

A render pass object represents a collection of attachments, subpasses,
and dependencies between the subpasses, and describes how the
attachments are used over the course of the subpasses.

Render passes are represented by `VkRenderPass` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkRenderPass)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html),
[VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html),
[VkSubpassShadingPipelineCreateInfoHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassShadingPipelineCreateInfoHUAWEI.html),
[vkCreateRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass.html),
[vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html),
[vkCreateRenderPass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2KHR.html),
[vkDestroyRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyRenderPass.html),
[vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceSubpassShadingMaxWorkgroupSizeHUAWEI.html),
[vkGetRenderAreaGranularity](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRenderAreaGranularity.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPass"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
