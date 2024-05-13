# VK_KHR_storage_buffer_storage_class(3) Manual Page

## Name

VK_KHR_storage_buffer_storage_class - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

132

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_storage_buffer_storage_class](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_storage_buffer_storage_class.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Alexander Galazin <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_storage_buffer_storage_class%5D%20@alegal-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_storage_buffer_storage_class%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>alegal-arm</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**IP Status**  
No known IP claims.

**Contributors**  
- Alexander Galazin, ARM

- David Neto, Google

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_KHR_storage_buffer_storage_class`

This extension provides a new SPIR-V `StorageBuffer` storage class. A
`Block`-decorated object in this class is equivalent to a
`BufferBlock`-decorated object in the `Uniform` storage class.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_STORAGE_BUFFER_STORAGE_CLASS_EXTENSION_NAME`

- `VK_KHR_STORAGE_BUFFER_STORAGE_CLASS_SPEC_VERSION`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-03-23 (Alexander Galazin)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_storage_buffer_storage_class"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
