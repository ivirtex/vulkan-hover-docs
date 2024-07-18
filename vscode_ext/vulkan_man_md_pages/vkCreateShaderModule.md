# vkCreateShaderModule(3) Manual Page

## Name

vkCreateShaderModule - Creates a new shader module object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a shader module, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateShaderModule(
    VkDevice                                    device,
    const VkShaderModuleCreateInfo*             pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkShaderModule*                             pShaderModule);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the shader module.

- `pCreateInfo` is a pointer to a
  [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structure.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pShaderModule` is a pointer to a
  [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) handle in which the resulting
  shader module object is returned.

## <a href="#_description" class="anchor"></a>Description

Once a shader module has been created, any entry points it contains
**can** be used in pipeline shader stages as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-compute"
target="_blank" rel="noopener">Compute Pipelines</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics"
target="_blank" rel="noopener">Graphics Pipelines</a>.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
target="_blank" rel="noopener"><code>maintenance5</code></a> feature is
enabled, shader module creation can be omitted entirely. Instead,
applications should provide the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html">VkShaderModuleCreateInfo</a>
structure directly in to pipeline creation by chaining it to <a
href="VkPipelineShaderStageCreateInfo.html">VkPipelineShaderStageCreateInfo</a>.
This avoids the overhead of creating and managing an additional
object.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCreateShaderModule-pCreateInfo-06904"
  id="VUID-vkCreateShaderModule-pCreateInfo-06904"></a>
  VUID-vkCreateShaderModule-pCreateInfo-06904  
  If `pCreateInfo` is not `NULL`, `pCreateInfo->pNext` **must** be
  `NULL` or a pointer to a valid instance of

  - [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html)

  - [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeaturesEXT.html)

Valid Usage (Implicit)

- <a href="#VUID-vkCreateShaderModule-device-parameter"
  id="VUID-vkCreateShaderModule-device-parameter"></a>
  VUID-vkCreateShaderModule-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateShaderModule-pCreateInfo-parameter"
  id="VUID-vkCreateShaderModule-pCreateInfo-parameter"></a>
  VUID-vkCreateShaderModule-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structure

- <a href="#VUID-vkCreateShaderModule-pAllocator-parameter"
  id="VUID-vkCreateShaderModule-pAllocator-parameter"></a>
  VUID-vkCreateShaderModule-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateShaderModule-pShaderModule-parameter"
  id="VUID-vkCreateShaderModule-pShaderModule-parameter"></a>
  VUID-vkCreateShaderModule-pShaderModule-parameter  
  `pShaderModule` **must** be a valid pointer to a
  [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_SHADER_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html),
[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateShaderModule"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
