# VkPhysicalDevicePipelineRobustnessFeaturesEXT(3) Manual Page

## Name

VkPhysicalDevicePipelineRobustnessFeaturesEXT - Structure describing
whether an implementation supports robustness requests on a per-pipeline
stage granularity



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDevicePipelineRobustnessFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_pipeline_robustness
typedef struct VkPhysicalDevicePipelineRobustnessFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           pipelineRobustness;
} VkPhysicalDevicePipelineRobustnessFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-pipelineRobustness"></span> `pipelineRobustness`
  indicates that robustness **can** be requested on a per-pipeline-stage
  granularity.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Enabling <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineRobustness"
target="_blank" rel="noopener"><code>pipelineRobustness</code></a> may,
on some platforms, incur a minor performance cost when <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
target="_blank" rel="noopener"><code>robustBufferAccess</code></a> is
disabled, even for pipelines which do not make use of any robustness
features. If robustness is not needed, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineRobustness"
target="_blank" rel="noopener"><code>pipelineRobustness</code></a>
should not be enabled by an application.</p></td>
</tr>
</tbody>
</table>

If the `VkPhysicalDevicePipelineRobustnessFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevicePipelineRobustnessFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDevicePipelineRobustnessFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDevicePipelineRobustnessFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDevicePipelineRobustnessFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_ROBUSTNESS_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_robustness](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_robustness.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevicePipelineRobustnessFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
