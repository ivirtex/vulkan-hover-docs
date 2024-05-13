# vkCompileDeferredNV(3) Manual Page

## Name

vkCompileDeferredNV - Deferred compilation of shaders



## <a href="#_c_specification" class="anchor"></a>C Specification

To compile a deferred shader in a pipeline call:

``` c
// Provided by VK_NV_ray_tracing
VkResult vkCompileDeferredNV(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    uint32_t                                    shader);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device containing the ray tracing pipeline.

- `pipeline` is the ray tracing pipeline object containing the shaders.

- `shader` is the index of the shader to compile.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCompileDeferredNV-pipeline-04621"
  id="VUID-vkCompileDeferredNV-pipeline-04621"></a>
  VUID-vkCompileDeferredNV-pipeline-04621  
  `pipeline` **must** be a ray tracing pipeline

- <a href="#VUID-vkCompileDeferredNV-pipeline-02237"
  id="VUID-vkCompileDeferredNV-pipeline-02237"></a>
  VUID-vkCompileDeferredNV-pipeline-02237  
  `pipeline` **must** have been created with
  `VK_PIPELINE_CREATE_DEFER_COMPILE_BIT_NV`

- <a href="#VUID-vkCompileDeferredNV-shader-02238"
  id="VUID-vkCompileDeferredNV-shader-02238"></a>
  VUID-vkCompileDeferredNV-shader-02238  
  `shader` **must** not have been called as a deferred compile before

Valid Usage (Implicit)

- <a href="#VUID-vkCompileDeferredNV-device-parameter"
  id="VUID-vkCompileDeferredNV-device-parameter"></a>
  VUID-vkCompileDeferredNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCompileDeferredNV-pipeline-parameter"
  id="VUID-vkCompileDeferredNV-pipeline-parameter"></a>
  VUID-vkCompileDeferredNV-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a href="#VUID-vkCompileDeferredNV-pipeline-parent"
  id="VUID-vkCompileDeferredNV-pipeline-parent"></a>
  VUID-vkCompileDeferredNV-pipeline-parent  
  `pipeline` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCompileDeferredNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
