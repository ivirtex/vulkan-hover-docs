# vkCreateRenderPass(3) Manual Page

## Name

vkCreateRenderPass - Create a new render pass object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a render pass, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateRenderPass(
    VkDevice                                    device,
    const VkRenderPassCreateInfo*               pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkRenderPass*                               pRenderPass);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the render pass.

- `pCreateInfo` is a pointer to a
  [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) structure
  describing the parameters of the render pass.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pRenderPass` is a pointer to a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle in which the resulting render pass object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreateRenderPass-device-10000"
  id="VUID-vkCreateRenderPass-device-10000"></a>
  VUID-vkCreateRenderPass-device-10000  
  `device` **must** support at least one queue family with the
  `VK_QUEUE_GRAPHICS_BIT` capability

Valid Usage (Implicit)

- <a href="#VUID-vkCreateRenderPass-device-parameter"
  id="VUID-vkCreateRenderPass-device-parameter"></a>
  VUID-vkCreateRenderPass-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateRenderPass-pCreateInfo-parameter"
  id="VUID-vkCreateRenderPass-pCreateInfo-parameter"></a>
  VUID-vkCreateRenderPass-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) structure

- <a href="#VUID-vkCreateRenderPass-pAllocator-parameter"
  id="VUID-vkCreateRenderPass-pAllocator-parameter"></a>
  VUID-vkCreateRenderPass-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateRenderPass-pRenderPass-parameter"
  id="VUID-vkCreateRenderPass-pRenderPass-parameter"></a>
  VUID-vkCreateRenderPass-pRenderPass-parameter  
  `pRenderPass` **must** be a valid pointer to a
  [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html),
[VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateRenderPass"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
