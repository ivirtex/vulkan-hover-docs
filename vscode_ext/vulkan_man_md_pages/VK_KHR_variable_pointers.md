# VK_KHR_variable_pointers(3) Manual Page

## Name

VK_KHR_variable_pointers - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

121

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
    
[VK_KHR_storage_buffer_storage_class](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_storage_buffer_storage_class.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_variable_pointers](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_variable_pointers.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_variable_pointers%5D%20@critsec%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_variable_pointers%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>critsec</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**IP Status**  
No known IP claims.

**Contributors**  
- John Kessenich, Google

- Neil Henning, Codeplay

- David Neto, Google

- Daniel Koch, Nvidia

- Graeme Leese, Broadcom

- Weifeng Zhang, Qualcomm

- Stephen Clarke, Imagination Technologies

- Faith Ekstrand, Intel

- Jesse Hall, Google

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_variable_pointers` extension allows implementations to
indicate their level of support for the `SPV_KHR_variable_pointers`
SPIR-V extension. The SPIR-V extension allows shader modules to use
invocation-private pointers into uniform and/or storage buffers, where
the pointer values can be dynamic and non-uniform.

The `SPV_KHR_variable_pointers` extension introduces two capabilities.
The first, `VariablePointersStorageBuffer`, **must** be supported by all
implementations of this extension. The second, `VariablePointers`, is
optional.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted, however support for the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-variablePointersStorageBuffer"
target="_blank"
rel="noopener"><code>variablePointersStorageBuffer</code></a> feature is
made optional. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceVariablePointerFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVariablePointerFeaturesKHR.html)

  - [VkPhysicalDeviceVariablePointersFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVariablePointersFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VARIABLE_POINTERS_EXTENSION_NAME`

- `VK_KHR_VARIABLE_POINTERS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-VariablePointers"
  target="_blank" rel="noopener"><code>VariablePointers</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-VariablePointersStorageBuffer"
  target="_blank"
  rel="noopener"><code>VariablePointersStorageBuffer</code></a>

## <a href="#_issues" class="anchor"></a>Issues

1\) Do we need an optional property for the SPIR-V
`VariablePointersStorageBuffer` capability or should it be mandatory
when this extension is advertised?

**RESOLVED**: Add it as a distinct feature, but make support mandatory.
Adding it as a feature makes the extension easier to include in a future
core API version. In the extension, the feature is mandatory, so that
presence of the extension guarantees some functionality. When included
in a core API version, the feature would be optional.

2\) Can support for these capabilities vary between shader stages?

**RESOLVED**: No, if the capability is supported in any stage it must be
supported in all stages.

3\) Should the capabilities be features or limits?

**RESOLVED**: Features, primarily for consistency with other similar
extensions.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-03-14 (Jesse Hall and John Kessenich)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_variable_pointers"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
