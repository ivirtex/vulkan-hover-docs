# vkGetPipelinePropertiesEXT(3) Manual Page

## Name

vkGetPipelinePropertiesEXT - Query pipeline properties



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the pipeline properties call:

``` c
// Provided by VK_EXT_pipeline_properties
VkResult vkGetPipelinePropertiesEXT(
    VkDevice                                    device,
    const VkPipelineInfoEXT*                    pPipelineInfo,
    VkBaseOutStructure*                         pPipelineProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the pipeline.

- `pPipelineInfo` is a pointer to a
  [VkPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInfoEXT.html) structure which describes
  the pipeline being queried.

- `pPipelineProperties` is a pointer to a
  [VkBaseOutStructure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBaseOutStructure.html) structure in which the
  pipeline properties will be written.

## <a href="#_description" class="anchor"></a>Description

To query a pipeline’s `pipelineIdentifier` pass a
[VkPipelinePropertiesIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelinePropertiesIdentifierEXT.html)
structure in `pPipelineProperties`. Each pipeline is associated with a
`pipelineIdentifier` and the identifier is implementation specific.

Valid Usage

- <a href="#VUID-vkGetPipelinePropertiesEXT-pipeline-06738"
  id="VUID-vkGetPipelinePropertiesEXT-pipeline-06738"></a>
  VUID-vkGetPipelinePropertiesEXT-pipeline-06738  
  The `pipeline` member of `pPipelineInfo` **must** have been created
  with `device`

- <a href="#VUID-vkGetPipelinePropertiesEXT-pPipelineProperties-06739"
  id="VUID-vkGetPipelinePropertiesEXT-pPipelineProperties-06739"></a>
  VUID-vkGetPipelinePropertiesEXT-pPipelineProperties-06739  
  `pPipelineProperties` **must** be a valid pointer to a
  [VkPipelinePropertiesIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelinePropertiesIdentifierEXT.html)
  structure

- <a href="#VUID-vkGetPipelinePropertiesEXT-None-06766"
  id="VUID-vkGetPipelinePropertiesEXT-None-06766"></a>
  VUID-vkGetPipelinePropertiesEXT-None-06766  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelinePropertiesIdentifier"
  target="_blank"
  rel="noopener"><code>pipelinePropertiesIdentifier</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkGetPipelinePropertiesEXT-device-parameter"
  id="VUID-vkGetPipelinePropertiesEXT-device-parameter"></a>
  VUID-vkGetPipelinePropertiesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetPipelinePropertiesEXT-pPipelineInfo-parameter"
  id="VUID-vkGetPipelinePropertiesEXT-pPipelineInfo-parameter"></a>
  VUID-vkGetPipelinePropertiesEXT-pPipelineInfo-parameter  
  `pPipelineInfo` **must** be a valid pointer to a valid
  [VkPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInfoEXT.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_properties.html),
[VkBaseOutStructure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBaseOutStructure.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPipelinePropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
