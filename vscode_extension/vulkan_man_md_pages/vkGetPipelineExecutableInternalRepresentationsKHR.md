# vkGetPipelineExecutableInternalRepresentationsKHR(3) Manual Page

## Name

vkGetPipelineExecutableInternalRepresentationsKHR - Get internal
representations of the pipeline executable



## <a href="#_c_specification" class="anchor"></a>C Specification

Each pipeline executable **may** have one or more text or binary
internal representations associated with it which are generated as part
of the compile process. These **may** include the final shader assembly,
a binary form of the compiled shader, or the shader compiler’s internal
representation at any number of intermediate compile steps. To query the
internal representations associated with a pipeline executable, call:

``` c
// Provided by VK_KHR_pipeline_executable_properties
VkResult vkGetPipelineExecutableInternalRepresentationsKHR(
    VkDevice                                    device,
    const VkPipelineExecutableInfoKHR*          pExecutableInfo,
    uint32_t*                                   pInternalRepresentationCount,
    VkPipelineExecutableInternalRepresentationKHR* pInternalRepresentations);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that created the pipeline.

- `pExecutableInfo` describes the pipeline executable being queried.

- `pInternalRepresentationCount` is a pointer to an integer related to
  the number of internal representations available or queried, as
  described below.

- `pInternalRepresentations` is either `NULL` or a pointer to an array
  of
  [VkPipelineExecutableInternalRepresentationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableInternalRepresentationKHR.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

If `pInternalRepresentations` is `NULL`, then the number of internal
representations associated with the pipeline executable is returned in
`pInternalRepresentationCount`. Otherwise,
`pInternalRepresentationCount` **must** point to a variable set by the
application to the number of elements in the `pInternalRepresentations`
array, and on return the variable is overwritten with the number of
structures actually written to `pInternalRepresentations`. If
`pInternalRepresentationCount` is less than the number of internal
representations associated with the pipeline executable, at most
`pInternalRepresentationCount` structures will be written, and
`VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate
that not all the available representations were returned.

While the details of the internal representations remain
implementation-dependent, the implementation **should** order the
internal representations in the order in which they occur in the
compiled pipeline with the final shader assembly (if any) last.

Valid Usage

- <a
  href="#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipelineExecutableInfo-03276"
  id="VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipelineExecutableInfo-03276"></a>
  VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipelineExecutableInfo-03276  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineExecutableInfo"
  target="_blank" rel="noopener"><code>pipelineExecutableInfo</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03277"
  id="VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03277"></a>
  VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03277  
  The `pipeline` member of `pExecutableInfo` **must** have been created
  with `device`

- <a
  href="#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03278"
  id="VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03278"></a>
  VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pipeline-03278  
  The `pipeline` member of `pExecutableInfo` **must** have been created
  with `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-device-parameter"
  id="VUID-vkGetPipelineExecutableInternalRepresentationsKHR-device-parameter"></a>
  VUID-vkGetPipelineExecutableInternalRepresentationsKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pExecutableInfo-parameter"
  id="VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pExecutableInfo-parameter"></a>
  VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pExecutableInfo-parameter  
  `pExecutableInfo` **must** be a valid pointer to a valid
  [VkPipelineExecutableInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableInfoKHR.html)
  structure

- <a
  href="#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentationCount-parameter"
  id="VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentationCount-parameter"></a>
  VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentationCount-parameter  
  `pInternalRepresentationCount` **must** be a valid pointer to a
  `uint32_t` value

- <a
  href="#VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentations-parameter"
  id="VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentations-parameter"></a>
  VUID-vkGetPipelineExecutableInternalRepresentationsKHR-pInternalRepresentations-parameter  
  If the value referenced by `pInternalRepresentationCount` is not `0`,
  and `pInternalRepresentations` is not `NULL`,
  `pInternalRepresentations` **must** be a valid pointer to an array of
  `pInternalRepresentationCount`
  [VkPipelineExecutableInternalRepresentationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableInternalRepresentationKHR.html)
  structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_pipeline_executable_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_executable_properties.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPipelineExecutableInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableInfoKHR.html),
[VkPipelineExecutableInternalRepresentationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableInternalRepresentationKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPipelineExecutableInternalRepresentationsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
