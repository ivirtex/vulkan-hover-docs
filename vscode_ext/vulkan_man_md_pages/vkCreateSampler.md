# vkCreateSampler(3) Manual Page

## Name

vkCreateSampler - Create a new sampler object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a sampler object, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateSampler(
    VkDevice                                    device,
    const VkSamplerCreateInfo*                  pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSampler*                                  pSampler);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the sampler.

- `pCreateInfo` is a pointer to a
  [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html) structure specifying
  the state of the sampler object.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pSampler` is a pointer to a [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) handle in
  which the resulting sampler object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreateSampler-device-09668"
  id="VUID-vkCreateSampler-device-09668"></a>
  VUID-vkCreateSampler-device-09668  
  `device` **must** support at least one queue family with one of the
  `VK_QUEUE_COMPUTE_BIT` or `VK_QUEUE_GRAPHICS_BIT` capabilities

- <a href="#VUID-vkCreateSampler-maxSamplerAllocationCount-04110"
  id="VUID-vkCreateSampler-maxSamplerAllocationCount-04110"></a>
  VUID-vkCreateSampler-maxSamplerAllocationCount-04110  
  There **must** be less than
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxSamplerAllocationCount`
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) objects currently created on the device

Valid Usage (Implicit)

- <a href="#VUID-vkCreateSampler-device-parameter"
  id="VUID-vkCreateSampler-device-parameter"></a>
  VUID-vkCreateSampler-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateSampler-pCreateInfo-parameter"
  id="VUID-vkCreateSampler-pCreateInfo-parameter"></a>
  VUID-vkCreateSampler-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html) structure

- <a href="#VUID-vkCreateSampler-pAllocator-parameter"
  id="VUID-vkCreateSampler-pAllocator-parameter"></a>
  VUID-vkCreateSampler-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateSampler-pSampler-parameter"
  id="VUID-vkCreateSampler-pSampler-parameter"></a>
  VUID-vkCreateSampler-pSampler-parameter  
  `pSampler` **must** be a valid pointer to a
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html),
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateSampler"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
