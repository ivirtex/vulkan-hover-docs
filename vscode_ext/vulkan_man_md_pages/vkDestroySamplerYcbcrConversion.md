# vkDestroySamplerYcbcrConversion(3) Manual Page

## Name

vkDestroySamplerYcbcrConversion - Destroy a created
Y′C<sub>B</sub>C<sub>R</sub> conversion



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a sampler Y′C<sub>B</sub>C<sub>R</sub> conversion, call:

``` c
// Provided by VK_VERSION_1_1
void vkDestroySamplerYcbcrConversion(
    VkDevice                                    device,
    VkSamplerYcbcrConversion                    ycbcrConversion,
    const VkAllocationCallbacks*                pAllocator);
```

or the equivalent command

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
void vkDestroySamplerYcbcrConversionKHR(
    VkDevice                                    device,
    VkSamplerYcbcrConversion                    ycbcrConversion,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the
  Y′C<sub>B</sub>C<sub>R</sub> conversion.

- `ycbcrConversion` is the conversion to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkDestroySamplerYcbcrConversion-device-parameter"
  id="VUID-vkDestroySamplerYcbcrConversion-device-parameter"></a>
  VUID-vkDestroySamplerYcbcrConversion-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parameter"
  id="VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parameter"></a>
  VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parameter  
  If `ycbcrConversion` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `ycbcrConversion` **must** be a valid
  [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) handle

- <a href="#VUID-vkDestroySamplerYcbcrConversion-pAllocator-parameter"
  id="VUID-vkDestroySamplerYcbcrConversion-pAllocator-parameter"></a>
  VUID-vkDestroySamplerYcbcrConversion-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parent"
  id="VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parent"></a>
  VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parent  
  If `ycbcrConversion` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `ycbcrConversion` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroySamplerYcbcrConversion"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
