# VkPhysicalDeviceSubgroupSizeControlFeatures(3) Manual Page

## Name

VkPhysicalDeviceSubgroupSizeControlFeatures - Structure describing the
subgroup size control features that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSubgroupSizeControlFeatures` structure is defined
as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceSubgroupSizeControlFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           subgroupSizeControl;
    VkBool32           computeFullSubgroups;
} VkPhysicalDeviceSubgroupSizeControlFeatures;
```

or the equivalent

``` c
// Provided by VK_EXT_subgroup_size_control
typedef VkPhysicalDeviceSubgroupSizeControlFeatures VkPhysicalDeviceSubgroupSizeControlFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-subgroupSizeControl"></span>
  `subgroupSizeControl` indicates whether the implementation supports
  controlling shader subgroup sizes via the
  `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT` flag
  and the
  [VkPipelineShaderStageRequiredSubgroupSizeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfo.html)
  structure.

- <span id="extension-features-computeFullSubgroups"></span>
  `computeFullSubgroups` indicates whether the implementation supports
  requiring full subgroups in compute , mesh, or task shaders via the
  `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` flag.

If the `VkPhysicalDeviceSubgroupSizeControlFeatures` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceSubgroupSizeControlFeatures` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>VkPhysicalDeviceSubgroupSizeControlFeaturesEXT</code>
structure was added in version 2 of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subgroup_size_control.html"><code>VK_EXT_subgroup_size_control</code></a>
extension. Version 1 implementations of this extension will not fill out
the features structure but applications may assume that both
<code>subgroupSizeControl</code> and <code>computeFullSubgroups</code>
are supported if the extension is supported. (See also the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-requirements"
target="_blank" rel="noopener">Feature Requirements</a> section.)
Applications are advised to add a
<code>VkPhysicalDeviceSubgroupSizeControlFeaturesEXT</code> structure to
the <code>pNext</code> chain of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html">VkDeviceCreateInfo</a> to enable the
features regardless of the version of the extension supported by the
implementation. If the implementation only supports version 1, it will
safely ignore the
<code>VkPhysicalDeviceSubgroupSizeControlFeaturesEXT</code>
structure.</p>
<p>Vulkan 1.3 implementations always support the features
structure.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceSubgroupSizeControlFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceSubgroupSizeControlFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceSubgroupSizeControlFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_subgroup_size_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subgroup_size_control.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSubgroupSizeControlFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
