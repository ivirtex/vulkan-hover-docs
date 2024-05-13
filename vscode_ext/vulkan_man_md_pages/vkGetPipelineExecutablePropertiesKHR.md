# vkGetPipelineExecutablePropertiesKHR(3) Manual Page

## Name

vkGetPipelineExecutablePropertiesKHR - Get the executables associated
with a pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

When a pipeline is created, its state and shaders are compiled into zero
or more device-specific executables, which are used when executing
commands against that pipeline. To query the properties of these
pipeline executables, call:

``` c
// Provided by VK_KHR_pipeline_executable_properties
VkResult vkGetPipelineExecutablePropertiesKHR(
    VkDevice                                    device,
    const VkPipelineInfoKHR*                    pPipelineInfo,
    uint32_t*                                   pExecutableCount,
    VkPipelineExecutablePropertiesKHR*          pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that created the pipeline.

- `pPipelineInfo` describes the pipeline being queried.

- `pExecutableCount` is a pointer to an integer related to the number of
  pipeline executables available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutablePropertiesKHR.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

If `pProperties` is `NULL`, then the number of pipeline executables
associated with the pipeline is returned in `pExecutableCount`.
Otherwise, `pExecutableCount` **must** point to a variable set by the
user to the number of elements in the `pProperties` array, and on return
the variable is overwritten with the number of structures actually
written to `pProperties`. If `pExecutableCount` is less than the number
of pipeline executables associated with the pipeline, at most
`pExecutableCount` structures will be written, and `VK_INCOMPLETE` will
be returned instead of `VK_SUCCESS`, to indicate that not all the
available properties were returned.

Valid Usage

- <a
  href="#VUID-vkGetPipelineExecutablePropertiesKHR-pipelineExecutableInfo-03270"
  id="VUID-vkGetPipelineExecutablePropertiesKHR-pipelineExecutableInfo-03270"></a>
  VUID-vkGetPipelineExecutablePropertiesKHR-pipelineExecutableInfo-03270  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineExecutableInfo"
  target="_blank" rel="noopener"><code>pipelineExecutableInfo</code></a>
  feature **must** be enabled

- <a href="#VUID-vkGetPipelineExecutablePropertiesKHR-pipeline-03271"
  id="VUID-vkGetPipelineExecutablePropertiesKHR-pipeline-03271"></a>
  VUID-vkGetPipelineExecutablePropertiesKHR-pipeline-03271  
  The `pipeline` member of `pPipelineInfo` **must** have been created
  with `device`

Valid Usage (Implicit)

- <a href="#VUID-vkGetPipelineExecutablePropertiesKHR-device-parameter"
  id="VUID-vkGetPipelineExecutablePropertiesKHR-device-parameter"></a>
  VUID-vkGetPipelineExecutablePropertiesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetPipelineExecutablePropertiesKHR-pPipelineInfo-parameter"
  id="VUID-vkGetPipelineExecutablePropertiesKHR-pPipelineInfo-parameter"></a>
  VUID-vkGetPipelineExecutablePropertiesKHR-pPipelineInfo-parameter  
  `pPipelineInfo` **must** be a valid pointer to a valid
  [VkPipelineInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInfoKHR.html) structure

- <a
  href="#VUID-vkGetPipelineExecutablePropertiesKHR-pExecutableCount-parameter"
  id="VUID-vkGetPipelineExecutablePropertiesKHR-pExecutableCount-parameter"></a>
  VUID-vkGetPipelineExecutablePropertiesKHR-pExecutableCount-parameter  
  `pExecutableCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPipelineExecutablePropertiesKHR-pProperties-parameter"
  id="VUID-vkGetPipelineExecutablePropertiesKHR-pProperties-parameter"></a>
  VUID-vkGetPipelineExecutablePropertiesKHR-pProperties-parameter  
  If the value referenced by `pExecutableCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pExecutableCount`
  [VkPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutablePropertiesKHR.html)
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
[VkPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutablePropertiesKHR.html),
[VkPipelineInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPipelineExecutablePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
