# vkCreateSamplerYcbcrConversion(3) Manual Page

## Name

vkCreateSamplerYcbcrConversion - Create a new
Y′C<sub>B</sub>C<sub>R</sub> conversion



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html),
call:

``` c
// Provided by VK_VERSION_1_1
VkResult vkCreateSamplerYcbcrConversion(
    VkDevice                                    device,
    const VkSamplerYcbcrConversionCreateInfo*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSamplerYcbcrConversion*                   pYcbcrConversion);
```

or the equivalent command

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
VkResult vkCreateSamplerYcbcrConversionKHR(
    VkDevice                                    device,
    const VkSamplerYcbcrConversionCreateInfo*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSamplerYcbcrConversion*                   pYcbcrConversion);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion.

- `pCreateInfo` is a pointer to a
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)
  structure specifying the requested sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pYcbcrConversion` is a pointer to a
  [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) handle in
  which the resulting sampler Y′C<sub>B</sub>C<sub>R</sub> conversion is
  returned.

## <a href="#_description" class="anchor"></a>Description

The interpretation of the configured sampler
Y′C<sub>B</sub>C<sub>R</sub> conversion is described in more detail in
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-sampler-YCbCr-conversion"
target="_blank" rel="noopener">the description of sampler
Y′C<sub>B</sub>C<sub>R</sub> conversion</a> in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures"
target="_blank" rel="noopener">Image Operations</a> chapter.

Valid Usage

- <a href="#VUID-vkCreateSamplerYcbcrConversion-None-01648"
  id="VUID-vkCreateSamplerYcbcrConversion-None-01648"></a>
  VUID-vkCreateSamplerYcbcrConversion-None-01648  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-samplerYcbcrConversion"
  target="_blank" rel="noopener"><code>samplerYcbcrConversion</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCreateSamplerYcbcrConversion-device-parameter"
  id="VUID-vkCreateSamplerYcbcrConversion-device-parameter"></a>
  VUID-vkCreateSamplerYcbcrConversion-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateSamplerYcbcrConversion-pCreateInfo-parameter"
  id="VUID-vkCreateSamplerYcbcrConversion-pCreateInfo-parameter"></a>
  VUID-vkCreateSamplerYcbcrConversion-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)
  structure

- <a href="#VUID-vkCreateSamplerYcbcrConversion-pAllocator-parameter"
  id="VUID-vkCreateSamplerYcbcrConversion-pAllocator-parameter"></a>
  VUID-vkCreateSamplerYcbcrConversion-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkCreateSamplerYcbcrConversion-pYcbcrConversion-parameter"
  id="VUID-vkCreateSamplerYcbcrConversion-pYcbcrConversion-parameter"></a>
  VUID-vkCreateSamplerYcbcrConversion-pYcbcrConversion-parameter  
  `pYcbcrConversion` **must** be a valid pointer to a
  [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html),
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateSamplerYcbcrConversion"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
