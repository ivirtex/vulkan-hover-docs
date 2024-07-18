# vkCreateRenderPass2(3) Manual Page

## Name

vkCreateRenderPass2 - Create a new render pass object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a render pass, call:

``` c
// Provided by VK_VERSION_1_2
VkResult vkCreateRenderPass2(
    VkDevice                                    device,
    const VkRenderPassCreateInfo2*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkRenderPass*                               pRenderPass);
```

or the equivalent command

``` c
// Provided by VK_KHR_create_renderpass2
VkResult vkCreateRenderPass2KHR(
    VkDevice                                    device,
    const VkRenderPassCreateInfo2*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkRenderPass*                               pRenderPass);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the render pass.

- `pCreateInfo` is a pointer to a
  [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html) structure
  describing the parameters of the render pass.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pRenderPass` is a pointer to a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle in which the resulting render pass object is returned.

## <a href="#_description" class="anchor"></a>Description

This command is functionally identical to
[vkCreateRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass.html), but includes extensible
sub-structures that include `sType` and `pNext` parameters, allowing
them to be more easily extended.

Valid Usage

- <a href="#VUID-vkCreateRenderPass2-device-10001"
  id="VUID-vkCreateRenderPass2-device-10001"></a>
  VUID-vkCreateRenderPass2-device-10001  
  `device` **must** support at least one queue family with the
  `VK_QUEUE_GRAPHICS_BIT` capability

Valid Usage (Implicit)

- <a href="#VUID-vkCreateRenderPass2-device-parameter"
  id="VUID-vkCreateRenderPass2-device-parameter"></a>
  VUID-vkCreateRenderPass2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateRenderPass2-pCreateInfo-parameter"
  id="VUID-vkCreateRenderPass2-pCreateInfo-parameter"></a>
  VUID-vkCreateRenderPass2-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html) structure

- <a href="#VUID-vkCreateRenderPass2-pAllocator-parameter"
  id="VUID-vkCreateRenderPass2-pAllocator-parameter"></a>
  VUID-vkCreateRenderPass2-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateRenderPass2-pRenderPass-parameter"
  id="VUID-vkCreateRenderPass2-pRenderPass-parameter"></a>
  VUID-vkCreateRenderPass2-pRenderPass-parameter  
  `pRenderPass` **must** be a valid pointer to a
  [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html),
[VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateRenderPass2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
